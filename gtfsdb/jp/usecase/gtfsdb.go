package usecase

import (
	"database/sql"
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/jp/domain/model"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/jp/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/paulmach/orb"
	"gorm.io/datatypes"
	"log/slog"
	"math"
	"os"
	"path"
	"sort"
	"sync"
	"time"
)

type GtfsJpDbUseCase interface {
	GtfsDbUrl(options CmdOptions) (digest string, err error)
	GtfsDbFile(options CmdOptions) (test string, err error)
	recalculateShapes() error
	recalculateShapesGeom() error
	recalculateStopTimes() error
	recalculateStopTimesGeom() error
	createShapesDetail() error
	createShapesDetailGeom() error
	createShapeEx() error
	createShapeExGeom() error
	createTripGeom() error
}

type gtfsJpDbUseCase struct {
	fileManagerRepo     repository.FileManagerRepository
	gtfsJpRepo          repository.GtfsJpRepository
	gtfsJpGeomRepo      repository.GtfsJpGeomRepository
	tripRepo            repository.TripRepository
	tripGeomRepo        repository.TripGeomRepository
	shapeRepo           repository.ShapeRepository
	shapeGeomRepo       repository.ShapeGeomRepository
	shapeExRepo         repository.ShapeExRepository
	shapeExGeomRepo     repository.ShapeExGeomRepository
	shapeDetailRepo     repository.ShapeDetailRepository
	shapeDetailGeomRepo repository.ShapeDetailGeomRepository
	stopTimeRepo        repository.StopTimeRepository
}

func (g gtfsJpDbUseCase) GtfsDbUrl(options CmdOptions) (digest string, err error) {
	// tmpディレクトリを作成
	tmp := "tmp"
	if err = os.MkdirAll(tmp, 0755); err != nil {
		return "", err
	}
	// gtfsをダウンロード
	options.GtfsFile = path.Join(tmp, "gtfs.zip")
	if err = g.fileManagerRepo.Download(options.GtfsUrl, options.GtfsFile); err != nil {
		return "", err
	}
	if digest, err = g.GtfsDbFile(options); err != nil {
		return "", err
	}
	return digest, nil
}

func (g gtfsJpDbUseCase) GtfsDbFile(options CmdOptions) (digest string, err error) {
	// tmpディレクトリを作成
	tmp := "tmp"
	if err = os.MkdirAll(tmp, 0777); err != nil {
		return "", err
	}

	// tmpディレクトリの削除
	defer func(fileManagerRepo repository.FileManagerRepository, path string) {
		_ = fileManagerRepo.Remove(path)
	}(g.fileManagerRepo, tmp)

	// gtfsを解凍
	gtfsPath, err := g.fileManagerRepo.UnZip(options.GtfsFile, tmp)
	if err != nil {
		return "", err
	}

	// digestの取得
	digest, err = util.Sha256(options.GtfsFile)
	if err != nil {
		return "", err
	}

	// option shapes_ex, shapes_detailがtrueの場合は、距離再計算が必須
	if options.ShapesDetail || options.ShapesEx {
		options.RecalculateDist = true
	}

	// マイグレーション, データ挿入
	if options.Geom {
		if err = g.gtfsJpGeomRepo.MigrateGtfsJpGeom(); err != nil {
			return "", err
		}
		if err = g.gtfsJpGeomRepo.CreateGtfsJpGeom(gtfsPath); err != nil {
			return "", err
		}
		if options.RecalculateDist {
			if err := g.recalculateShapesGeom(); err != nil {
				return "", err
			}
		}
		if options.ShapesDetail {
			if err := g.shapeDetailGeomRepo.MigrateShapesDetailGeom(); err != nil {
				return "", err
			}
			if err := g.createShapesDetailGeom(); err != nil {
				return "", err
			}
		}
		if options.ShapesEx {
			if err := g.shapeExGeomRepo.MigrateShapesExGeom(); err != nil {
				return "", err
			}
			if err := g.createShapeExGeom(); err != nil {
				return "", err
			}
		}
		//Geom(linestring)を更新する処理を実行する
		if err = g.createTripGeom(); err != nil {
			return "", err
		}
	} else {
		if err = g.gtfsJpRepo.MigrateGtfsJp(); err != nil {
			return "", err
		}
		if err = g.gtfsJpRepo.CreateGtfsJp(gtfsPath); err != nil {
			return "", err
		}
		if options.RecalculateDist {
			if err := g.recalculateShapes(); err != nil {
				return "", err
			}
		}
		if options.ShapesDetail {
			if err := g.shapeDetailRepo.MigrateShapesDetail(); err != nil {
				return "", err
			}
			if err := g.createShapesDetail(); err != nil {
				return "", err
			}
		}
		if options.ShapesEx {
			if err := g.shapeExRepo.MigrateShapesEx(); err != nil {
				return "", err
			}
			if err := g.createShapeEx(); err != nil {
				return "", err
			}
		}

	}

	return "", err
}

func (g gtfsJpDbUseCase) recalculateShapes() error {
	slog.Info("テーブル[shapes] shape_dist_traveled の再計算を実行します")

	// shape_idのスライスを取得
	shapeIds, err := g.shapeRepo.FindShapeIds()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 50)       // セマフォチャンネル
	errChan := make(chan error, len(shapeIds)) // エラーチャンネル

	for _, shapeId := range shapeIds {
		wg.Add(1)

		semaphore <- struct{}{} // セマフォを取得
		go func(shapeId string) {
			defer wg.Done()
			defer func() { <-semaphore }() // 処理終了後セマフォを解放

			var recalculateShapes []model.Shape
			totalDistance := 0.0

			shapes, err := g.shapeRepo.FindShapesByShapeId(shapeId)
			if err != nil {
				errChan <- err
				return
			}

			// 距離の再計算
			for i, pt := range shapes {
				if i == 0 {
					distTraveled := math.Floor(totalDistance)
					pt.ShapeDistTraveled = &distTraveled
					recalculateShapes = append(recalculateShapes, pt)
					continue
				}
				// 2点間の距離を計算
				totalDistance += util.KarneyWgs84(
					shapes[i-1].ShapePtLat, shapes[i-1].ShapePtLon,
					pt.ShapePtLat, pt.ShapePtLon,
				)
				distTraveled := math.Floor(totalDistance)
				pt.ShapeDistTraveled = &distTraveled
				recalculateShapes = append(recalculateShapes, pt)
			}

			// 再計算したデータをDBに格納
			if err := g.shapeRepo.UpdateShapes(recalculateShapes); err != nil {
				errChan <- err
				return
			}
		}(shapeId)
	}

	wg.Wait()
	close(errChan)

	// エラーがあれば最初のものを返す
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	slog.Info("テーブル[shapes] shape_dist_traveled の再計算完了")
	return nil
}

func (g gtfsJpDbUseCase) recalculateShapesGeom() error {
	slog.Info("テーブル[shapes] shape_dist_traveled の再計算を実行します")

	// shape_idのスライスを取得
	shapeIds, err := g.shapeGeomRepo.FindShapeGeomIds()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 50)       // セマフォチャンネル
	errChan := make(chan error, len(shapeIds)) // エラーチャンネル

	for _, shapeId := range shapeIds {
		wg.Add(1)

		semaphore <- struct{}{} // セマフォを取得
		go func(shapeId string) {
			defer wg.Done()
			defer func() { <-semaphore }() // 処理終了後セマフォを解放

			var recalculateShapes []model.ShapeGeom
			totalDistance := 0.0

			shapes, err := g.shapeGeomRepo.FindShapesGeomByShapeId(shapeId)
			if err != nil {
				errChan <- err
				return
			}

			// 距離の再計算
			for i, pt := range shapes {
				if i == 0 {
					distTraveled := math.Floor(totalDistance)
					pt.ShapeDistTraveled = &distTraveled
					recalculateShapes = append(recalculateShapes, pt)
					continue
				}
				// 2点間の距離を計算
				totalDistance += util.KarneyWgs84(
					shapes[i-1].ShapePtLat, shapes[i-1].ShapePtLon,
					pt.ShapePtLat, pt.ShapePtLon,
				)
				distTraveled := math.Floor(totalDistance)
				pt.ShapeDistTraveled = &distTraveled
				recalculateShapes = append(recalculateShapes, pt)
			}

			// 再計算したデータをDBに格納
			if err := g.shapeGeomRepo.UpdateShapesGeom(recalculateShapes); err != nil {
				errChan <- err
				return
			}
		}(shapeId)
	}

	wg.Wait()
	close(errChan)

	// エラーがあれば最初のものを返す
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	slog.Info("テーブル[shapes] shape_dist_traveled の再計算完了")
	return nil
}

func (g gtfsJpDbUseCase) recalculateStopTimes() error {
	//tripIds取得
	//tripIdごとにshape取得
	//tripIdごとにstopTimesWithStops取得(時間順)
	//shapeとstopTimesWithStopsの位置が一番近いものを検索
	//
	//TODO 保留中
	panic("stop_timesの距離再計算は保留中")
}

func (g gtfsJpDbUseCase) recalculateStopTimesGeom() error {
	//TODO 保留中
	panic("stop_timesの距離再計算は保留中")
}

func (g gtfsJpDbUseCase) createShapesDetail() error {
	slog.Info("テーブル[shapesDetail] 作成開始 数分かかる場合があります")

	// shapeの間隔
	interval := 5.0

	shapeIds, err := g.shapeRepo.FindShapeIds()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(shapeIds))
	semaphore := make(chan struct{}, 50) // 同時に実行されるゴルーチンの最大数を50に制限

	for _, shapeId := range shapeIds {
		wg.Add(1)
		semaphore <- struct{}{} // セマフォを取得
		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphore }() // ゴルーチンが終了したらセマフォを解放

			shapes, err := g.shapeRepo.FindShapesByShapeId(id)
			if err != nil {
				errChan <- err
				return
			}

			// 最初のshape_detailデータを追加
			shapesDetail := []model.ShapeDetail{{
				ShapeDetail: gtfsjp.ShapeDetail{
					ShapeId:               id,
					ShapePtLat:            shapes[0].ShapePtLat,
					ShapePtLon:            shapes[0].ShapePtLon,
					ShapeDetailPtSequence: shapes[0].ShapePtSequence,
					ShapeDistTraveled:     *shapes[0].ShapeDistTraveled,
				},
			}}

			shapePtSeqCnt := 0

			for i := 0; i < len(shapes)-1; i++ {
				currentShapePtLat := shapes[i].ShapePtLat
				currentShapePtLon := shapes[i].ShapePtLon
				nextShapePtLat := shapes[i+1].ShapePtLat
				nextShapePtLon := shapes[i+1].ShapePtLon

				blockDistance := util.KarneyWgs84(currentShapePtLat, currentShapePtLon, nextShapePtLat, nextShapePtLon)

				// shapeの点が重なっている場合、処理をせずに追加
				if blockDistance == 0 {
					var nextShapeDistTraveled float64
					nextShapeDistTraveled = shapesDetail[shapePtSeqCnt].ShapeDistTraveled + (blockDistance / interval)
					nextShapeDistTraveled = math.Round(nextShapeDistTraveled*10) / 10
					shapePtSeqCnt++

					shapesDetail = append(shapesDetail, model.ShapeDetail{
						ShapeDetail: gtfsjp.ShapeDetail{
							ShapeId:               id,
							ShapePtLat:            nextShapePtLat,
							ShapePtLon:            nextShapePtLon,
							ShapeDetailPtSequence: shapePtSeqCnt,
							ShapeDistTraveled:     nextShapeDistTraveled,
						},
					})
					continue
				}

				// 新規にshapeDetailデータを生成・追加
				dLat := nextShapePtLat - currentShapePtLat
				dLon := nextShapePtLon - currentShapePtLon

				for j := 1; j <= int(interval)-1; j++ {
					nextLat := (1/interval)*dLat + currentShapePtLat
					nextLon := (1/interval)*dLon + currentShapePtLon

					var nextShapeDistTraveled float64
					nextShapeDistTraveled = shapesDetail[shapePtSeqCnt].ShapeDistTraveled + (blockDistance / interval)
					nextShapeDistTraveled = math.Round(nextShapeDistTraveled*10) / 10
					shapePtSeqCnt++

					shapesDetail = append(
						shapesDetail, model.ShapeDetail{
							ShapeDetail: gtfsjp.ShapeDetail{
								ShapeId:               id,
								ShapePtLat:            nextLat,
								ShapePtLon:            nextLon,
								ShapeDetailPtSequence: shapePtSeqCnt,
								ShapeDistTraveled:     nextShapeDistTraveled,
							},
						},
					)
					currentShapePtLat = nextLat
					currentShapePtLon = nextLon
				}

				// 既存のshapeデータを追加
				nextShapeDistTraveled := shapesDetail[shapePtSeqCnt].ShapeDistTraveled + (blockDistance / interval)
				nextShapeDistTraveled = math.Round(nextShapeDistTraveled*10) / 10
				shapePtSeqCnt++

				shapesDetail = append(
					shapesDetail, model.ShapeDetail{
						ShapeDetail: gtfsjp.ShapeDetail{
							ShapeId:               id,
							ShapePtLat:            nextShapePtLat,
							ShapePtLon:            nextShapePtLon,
							ShapeDetailPtSequence: shapePtSeqCnt,
							ShapeDistTraveled:     nextShapeDistTraveled,
						},
					},
				)

			}

			if err := g.shapeDetailRepo.CreateShapesDetail(shapesDetail); err != nil {
				errChan <- err
				return
			}
		}(shapeId)
	}

	wg.Wait()
	close(errChan)
	if len(errChan) > 0 {
		return <-errChan // 最初のエラーを返す
	}

	slog.Info("テーブル[shapesDetail] 作成完了")
	return nil
}

func (g gtfsJpDbUseCase) createShapesDetailGeom() error {
	slog.Info("テーブル[shapesDetail] 作成開始 数分かかる場合があります")

	// shapeの間隔
	interval := 5.0

	shapeIds, err := g.shapeGeomRepo.FindShapeGeomIds()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(shapeIds))
	semaphore := make(chan struct{}, 50) // 同時に実行されるゴルーチンの最大数を50に制限

	for _, shapeId := range shapeIds {
		wg.Add(1)
		semaphore <- struct{}{} // セマフォを取得
		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphore }() // ゴルーチンが終了したらセマフォを解放

			shapes, err := g.shapeGeomRepo.FindShapesGeomByShapeId(id)
			if err != nil {
				errChan <- err
				return
			}

			shapesDetail := []model.ShapeDetailGeom{{
				ShapeDetailGeom: gtfsjp.ShapeDetailGeom{
					ShapeId:               id,
					ShapePtLat:            shapes[0].ShapePtLat,
					ShapePtLon:            shapes[0].ShapePtLon,
					ShapeDetailPtSequence: shapes[0].ShapePtSequence,
					ShapeDistTraveled:     *shapes[0].ShapeDistTraveled,
					Geom:                  shapes[0].Geom,
				},
			}}

			remainder := 0.0
			shapePtSeqCnt := shapes[0].ShapePtSequence

			for i := 1; i < len(shapes); i++ {
				prevShapePtLat := shapesDetail[i-1].ShapePtLat
				prevShapePtLon := shapesDetail[i-1].ShapePtLon
				nextShapePtLat := shapes[i].ShapePtLat
				nextShapePtLon := shapes[i].ShapePtLon

				blockDistance := util.KarneyWgs84(prevShapePtLat, prevShapePtLon, nextShapePtLat, nextShapePtLon)
				blockDistance += remainder

				repeat := int(blockDistance / interval)
				remainder = math.Mod(blockDistance, interval)

				if repeat == 0 {
					shapePtSeqCnt++
					shapesDetail = append(shapesDetail, model.ShapeDetailGeom{
						ShapeDetailGeom: gtfsjp.ShapeDetailGeom{
							ShapeId:               id,
							ShapePtLat:            shapes[i].ShapePtLat,
							ShapePtLon:            shapes[i].ShapePtLon,
							ShapeDetailPtSequence: shapePtSeqCnt,
							ShapeDistTraveled:     math.Round((shapesDetail[len(shapesDetail)-1].ShapeDistTraveled+remainder)*10) / 10,
							Geom:                  gormdatatypes.Geometry{Geom: orb.Point{shapes[i].ShapePtLon, shapes[i].ShapePtLat}, Srid: 4326},
						},
					})
					remainder = 0.0
					continue
				}

				t := interval / blockDistance
				dLat := nextShapePtLat - prevShapePtLat
				dLon := nextShapePtLon - prevShapePtLon

				for j := 0; j < repeat; j++ {
					shapePtSeqCnt++
					prevLat := shapesDetail[len(shapesDetail)-1].ShapePtLat
					prevLon := shapesDetail[len(shapesDetail)-1].ShapePtLon
					nextLat := t*dLat + prevLat
					nextLon := t*dLon + prevLon

					shortDistance := util.KarneyWgs84(prevLat, prevLon, nextLat, nextLon)

					shapeDistTraveled := shapesDetail[len(shapesDetail)-1].ShapeDistTraveled + shortDistance
					shapesDetail = append(shapesDetail, model.ShapeDetailGeom{
						ShapeDetailGeom: gtfsjp.ShapeDetailGeom{
							ShapeId:               id,
							ShapePtLat:            nextLat,
							ShapePtLon:            nextLon,
							ShapeDetailPtSequence: shapePtSeqCnt,
							ShapeDistTraveled:     math.Round(shapeDistTraveled*10) / 10,
							Geom:                  gormdatatypes.Geometry{Geom: orb.Point{nextLon, nextLat}, Srid: 4326},
						},
					})
				}
			}

			if err := g.shapeDetailGeomRepo.CreateShapesDetailGeom(shapesDetail); err != nil {
				errChan <- err
				return
			}
		}(shapeId)
	}

	wg.Wait()
	close(errChan)
	if len(errChan) > 0 {
		return <-errChan // 最初のエラーを返す
	}

	slog.Info("テーブル[shapesDetail] 作成完了")
	return nil
}

func (g gtfsJpDbUseCase) createShapeEx() error {
	slog.Info("テーブル[shapes_ex] 作成開始")
	// stops, stop_timesを結合して作成したshape_exの取得（stop_idの情報はない）
	tmpShapesEx, err := g.shapeExRepo.FindShapesExByTripsAndShapes() //trip_idに対応するshapeの情報がくっついたもの
	if err != nil {
		return err
	}

	if err := g.shapeExRepo.CreateShapesEx(tmpShapesEx); err != nil { // さっき作ったshape_exデータをDBに書き込んでいる
		return err
	}

	tripIds, err := g.tripRepo.FindTripIds() // trip_idの一覧を取得
	if err != nil {
		return err
	}

	for _, tripId := range tripIds {
		var insertShapesExTemp []model.ShapeExTemp
		var insertShapesEx []model.ShapeEx
		shapesEx, err := g.shapeExRepo.FindShapesExByTripId(tripId) // さっき作ったshapes_exのテーブルに対してtrip_idを指定してshapes_exデータを取得
		if err != nil {
			return err
		}
		if len(shapesEx) == 0 {
			slog.Warn(fmt.Sprintf("trip_id: %sのshapesが存在しないためスキップします", tripId))
			continue
		}

		stopLocationStr, err := g.shapeExRepo.FindTripWithStopLocationByTripId(tripId) //stop_timeにバス停のくっついたデータに対して指定のtrip_idに対応するデータを取得
		if err != nil {
			return err
		}

		var stopLocation []model.TripWithStopLocation
		for _, r := range stopLocationStr {
			arrivalTime, _ := time.Parse("15:04:05", r.Arrival)
			departure, _ := time.Parse("15:04:05", r.Departure)

			stopLocation = append(stopLocation, model.TripWithStopLocation{
				TripId:        r.TripId,
				StopId:        r.StopId,
				StopSequence:  r.StopSequence,
				StopLat:       r.StopLat,
				StopLon:       r.StopLon,
				ArrivalTime:   arrivalTime,
				DepartureTime: departure,
			})
		}

		for _, s := range shapesEx {
			insertShapesExTemp = append(insertShapesExTemp, model.ShapeExTemp{
				ShapeExTemp: gtfsjp.ShapeExTemp{
					TripId:          tripId,
					ShapeId:         s.ShapeId,
					ShapePtSequence: s.ShapePtSequence,
					ShapePtLat:      s.ShapePtLat,
					ShapePtLon:      s.ShapePtLon,
				},
			})
		}

		shapesLen := len(shapesEx)
		idx := 0

		for _, stop := range stopLocation { //shape_exの一行に対してバス停データを緯度経度の距離を見て総当たりで一番近いものを見つけてstop_id情報をくっつける
			minDist := math.MaxFloat64
			for i := idx; i < shapesLen; i++ {
				dist := util.KarneyWgs84(stop.StopLat, stop.StopLon, shapesEx[i].ShapePtLat, shapesEx[i].ShapePtLon)
				if dist < minDist {
					minDist = dist
					idx = i
				}
				if minDist <= 0 {
					break
				}
			}

			insertShapesExTemp[idx].StopId = sql.NullString{
				String: stop.StopId,
				Valid:  true,
			}

			insertShapesExTemp[idx].ShapesTime = stop.ArrivalTime
		}

		// ここまででshapes_exデータにarrival_timeとdeparture_timeカラムがくっついたデータができてる
		sort.Slice(insertShapesExTemp, func(i, j int) bool {
			return insertShapesExTemp[i].ShapePtSequence < insertShapesExTemp[j].ShapePtSequence
		})

		firstId := 0
		secondId := 0

		for i := 1; i < len(insertShapesExTemp); i++ { //insertは更新されたデータしか入っていないから今回の場合はstop_idが入ったデータだけ
			if insertShapesExTemp[i].StopId.Valid { //stop_idカラムに値が入っている
				secondId = i
				shapesNum := secondId - firstId

				if shapesNum == 1 {
					continue
				}

				temp := insertShapesExTemp[secondId].ShapesTime.Sub(insertShapesExTemp[firstId].ShapesTime)
				segment := temp / time.Duration(shapesNum)

				for j := firstId + 1; j < secondId; j++ {
					insertShapesExTemp[j].ShapesTime = insertShapesExTemp[j-1].ShapesTime.Add(segment)
				}
				firstId = secondId
			}
		}

		// 仮で作成したinsertShapesExTempからarrival_time、departure_time属性を除いたshape_ex構造体にデータ型を合わせる
		for i := 0; i < len(insertShapesExTemp); i++ {

			// time.Timeをdatatypes.Time に変換
			hour, minute, sec := insertShapesExTemp[i].ShapesTime.Clock()
			nsec := insertShapesExTemp[i].ShapesTime.Nanosecond()
			shapesTime := datatypes.NewTime(hour, minute, sec, nsec)

			insertShapesEx = append(insertShapesEx, model.ShapeEx{
				ShapeEx: gtfsjp.ShapeEx{
					TripId:          tripId,
					ShapeId:         insertShapesExTemp[i].ShapeId,
					ShapePtSequence: insertShapesExTemp[i].ShapePtSequence,
					StopId:          insertShapesExTemp[i].StopId,
					ShapesTime:      shapesTime,
				},
			})
		}

		if err := g.shapeExRepo.UpdateShapesEx(insertShapesEx); err != nil {
			return err
		}
	}
	slog.Info("テーブル[shapes_ex] 作成完了")
	return nil
}

func (g gtfsJpDbUseCase) createShapeExGeom() error {
	slog.Info("テーブル[shapes_ex] 作成開始")
	// stops, stop_timesを結合して作成したshape_exの取得（stop_idの情報はない）
	tmpShapesExGeom, err := g.shapeExGeomRepo.FindShapesExGeomByTripsAndShapes()
	if err != nil {
		return err
	}

	if err := g.shapeExGeomRepo.CreateShapesExGeom(tmpShapesExGeom); err != nil {
		return err
	}

	tripIds, err := g.tripRepo.FindTripIds()
	if err != nil {
		return err
	}

	for _, tripId := range tripIds {
		var insertShapesExGeom []model.ShapeExGeom
		shapesExGeom, err := g.shapeExGeomRepo.FindShapesExGeomByTripId(tripId)
		if err != nil {
			return err
		}
		if len(shapesExGeom) == 0 {
			slog.Warn(fmt.Sprintf("trip_id: %sのshapesが存在しないためスキップします", tripId))
			continue
		}

		stopLocation, err := g.shapeExGeomRepo.FindTripWithStopLocationByTripId(tripId)
		if err != nil {
			return err
		}

		shapesLen := len(shapesExGeom)
		idx := 0

		for _, stop := range stopLocation {
			minDist := math.MaxFloat64
			for i := idx; i < shapesLen; i++ {
				dist := util.KarneyWgs84(stop.StopLat, stop.StopLon, shapesExGeom[i].ShapePtLat, shapesExGeom[i].ShapePtLon)
				if dist < minDist {
					minDist = dist
					idx = i
				}
				if minDist <= 0 {
					break
				}
			}

			insertShapesExGeom = append(insertShapesExGeom, model.ShapeExGeom{
				ShapeExGeom: gtfsjp.ShapeExGeom{
					TripId:          tripId,
					ShapeId:         shapesExGeom[idx].ShapeId,
					ShapePtSequence: shapesExGeom[idx].ShapePtSequence,
					StopId: sql.NullString{
						String: stop.StopId,
						Valid:  true,
					},
				},
			})
		}
		if err := g.shapeExGeomRepo.UpdateShapesExGeom(insertShapesExGeom); err != nil {
			return err
		}
	}
	slog.Info("テーブル[shapes_ex] 作成完了")
	return nil
}

func (g gtfsJpDbUseCase) createTripGeom() error {
	slog.Info("テーブル[trip] geom の計算を実行します")
	//tripIdの一覧表を取得する
	tripIds, err := g.tripRepo.FindTripIds()
	if err != nil {
		return err
	}
	var tripGeom []model.TripGeomLine

	//tripIdごとにGeomlineを生成、更新用の構造体に格納していく
	for _, tripId := range tripIds {
		shapeId, err := g.tripRepo.FindShapeIdByTripId(tripId)
		if err != nil {
			return err
		}
		shapesGeom, err := g.shapeGeomRepo.FindShapesGeomByShapeId(shapeId)
		if err != nil {
			return nil
		}

		//shapesGeomのgeomフィールドのみのデータからなるスライスを作成
		var geoms []gormdatatypes.Geometry
		for _, shapeGeom := range shapesGeom {
			geoms = append(geoms, shapeGeom.Geom)
		}

		//gormdatatypes.Geometry型のスライスをorb.Point型のスライスに変換する
		var pointSlice []orb.Point
		for _, geom := range geoms {
			point := geom.Geom.(orb.Point)
			pointSlice = append(pointSlice, point)
		}

		//Geomのlinestring型のデータを生成する
		GeomLine := orb.LineString(pointSlice)

		//更新用の構造体にデータを入れる
		tripGeom = append(tripGeom, model.TripGeomLine{
			TripId: tripId,
			Geom:   gormdatatypes.Geometry{Geom: GeomLine, Srid: 4326},
		})
	}
	//tripテーブルのGeomカラムのデータをstring_line型で更新する
	if err := g.tripGeomRepo.UpdateTripsGeom(tripGeom); err != nil {
		return err
	}
	slog.Info("テーブル[trip] geomの計算完了")
	return nil
}

func NewGtfsJpDbUseCase(fileManagerRepository repository.FileManagerRepository,
	gtfsJpRepository repository.GtfsJpRepository,
	gtfsJpGeomRepository repository.GtfsJpGeomRepository,
	tripRepository repository.TripRepository,
	tripGeomRepository repository.TripGeomRepository,
	shapeRepository repository.ShapeRepository,
	shapeGeomRepository repository.ShapeGeomRepository,
	shapeExRepository repository.ShapeExRepository,
	shapeExGeomRepository repository.ShapeExGeomRepository,
	shapeDetailRepository repository.ShapeDetailRepository,
	shapeDetailGeomRepository repository.ShapeDetailGeomRepository,
	stopTimeRepository repository.StopTimeRepository) GtfsJpDbUseCase {
	return gtfsJpDbUseCase{
		fileManagerRepo:     fileManagerRepository,
		gtfsJpRepo:          gtfsJpRepository,
		gtfsJpGeomRepo:      gtfsJpGeomRepository,
		tripRepo:            tripRepository,
		tripGeomRepo:        tripGeomRepository,
		shapeRepo:           shapeRepository,
		shapeGeomRepo:       shapeGeomRepository,
		shapeExRepo:         shapeExRepository,
		shapeExGeomRepo:     shapeExGeomRepository,
		shapeDetailRepo:     shapeDetailRepository,
		shapeDetailGeomRepo: shapeDetailGeomRepository,
		stopTimeRepo:        stopTimeRepository,
	}
}

type CmdOptions struct {
	GtfsUrl         string
	GtfsFile        string
	ShapesEx        bool
	ShapesDetail    bool
	Geom            bool
	RecalculateDist bool
	Dsn             string
	Schema          string
}
