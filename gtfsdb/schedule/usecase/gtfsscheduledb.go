package usecase

import (
	"database/sql"
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/schedule/domain/model"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/schedule/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsschedule"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/paulmach/orb"
	"log/slog"
	"math"
	"os"
	"path"
	"sync"
)

type GtfsScheduleDbUseCase interface {
	GtfsDbUrl(options CmdOptions) (digest string, err error)
	GtfsDbFile(options CmdOptions) (digest string, err error)
	recalculateShapes() error
	recalculateShapesGeom() error
	recalculateStopTimes() error
	recalculateStopTimesGeom() error
	createShapesDetail() error
	createShapesDetailGeom() error
	createShapeEx() error
	createShapeExGeom() error
}

type gtfsScheduleDbUseCase struct {
	fileManagerRepo      repository.FileManagerRepository
	gtfsScheduleRepo     repository.GtfsScheduleRepository
	gtfsScheduleGeomRepo repository.GtfsScheduleGeomRepository
	tripRepo             repository.TripRepository
	tripGeomRepo         repository.TripGeomRepository
	shapeRepo            repository.ShapeRepository
	shapeGeomRepo        repository.ShapeGeomRepository
	shapeExRepo          repository.ShapeExRepository
	shapeExGeomRepo      repository.ShapeExGeomRepository
	shapeDetailRepo      repository.ShapeDetailRepository
	shapeDetailGeomRepo  repository.ShapeDetailGeomRepository
	stopTimeRepo         repository.StopTimeRepository
}

func (g gtfsScheduleDbUseCase) GtfsDbUrl(options CmdOptions) (digest string, err error) {
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

func (g gtfsScheduleDbUseCase) GtfsDbFile(options CmdOptions) (digest string, err error) {
	// tmpディレクトリを作成
	tmp := "tmp"
	if err = os.MkdirAll(tmp, 0755); err != nil {
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
		if err = g.gtfsScheduleGeomRepo.MigrateGtfsScheduleGeom(); err != nil {
			return "", err
		}
		if err = g.gtfsScheduleGeomRepo.CreateGtfsScheduleGeom(gtfsPath); err != nil {
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

	} else {
		if err = g.gtfsScheduleRepo.MigrateGtfsSchedule(); err != nil {
			return "", err
		}
		if err = g.gtfsScheduleRepo.CreateGtfsSchedule(gtfsPath); err != nil {
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

	return digest, err
}

func (g gtfsScheduleDbUseCase) recalculateShapes() error {
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

func (g gtfsScheduleDbUseCase) recalculateShapesGeom() error {
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

func (g gtfsScheduleDbUseCase) recalculateStopTimes() error {
	//tripIds取得
	//tripIdごとにshape取得
	//tripIdごとにstopTimesWithStops取得(時間順)
	//shapeとstopTimesWithStopsの位置が一番近いものを検索
	//
	//TODO 保留中
	panic("stop_timesの距離再計算は保留中")
}

func (g gtfsScheduleDbUseCase) recalculateStopTimesGeom() error {
	//TODO 保留中
	panic("stop_timesの距離再計算は保留中")
}

func (g gtfsScheduleDbUseCase) createShapesDetail() error {
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

			shapesDetail := []model.ShapeDetail{{
				ShapeDetail: gtfsschedule.ShapeDetail{
					ShapeId:               id,
					ShapePtLat:            shapes[0].ShapePtLat,
					ShapePtLon:            shapes[0].ShapePtLon,
					ShapeDetailPtSequence: shapes[0].ShapePtSequence,
					ShapeDistTraveled:     *shapes[0].ShapeDistTraveled,
				},
			}}

			shapePtSeqCnt := shapes[0].ShapePtSequence

			for i := 1; i < len(shapes); i++ {
				prevShapePtLat := shapes[i-1].ShapePtLat
				prevShapePtLon := shapes[i-1].ShapePtLon
				nextShapePtLat := shapes[i].ShapePtLat
				nextShapePtLon := shapes[i].ShapePtLon

				blockDistance := util.KarneyWgs84(prevShapePtLat, prevShapePtLon, nextShapePtLat, nextShapePtLon)
				repeat := int(blockDistance / interval)

				if repeat == 0 {
					nextShapeDistTraveled := math.Round((shapesDetail[len(shapesDetail)-1].ShapeDistTraveled)*10) / 10
					shapePtSeqCnt++
					shapesDetail = append(shapesDetail, model.ShapeDetail{
						ShapeDetail: gtfsschedule.ShapeDetail{
							ShapeId:               id,
							ShapePtLat:            nextShapePtLat,
							ShapePtLon:            nextShapePtLon,
							ShapeDetailPtSequence: shapePtSeqCnt,
							ShapeDistTraveled:     nextShapeDistTraveled,
						},
					})
					continue
				}

				t := interval / blockDistance
				dLat := nextShapePtLat - prevShapePtLat
				dLon := nextShapePtLon - prevShapePtLon

				for j := 1; j <= repeat; j++ {
					nextLat := float64(j)*t*dLat + prevShapePtLat
					nextLon := float64(j)*t*dLon + prevShapePtLon

					nextShapeDistTraveled := shapesDetail[len(shapesDetail)-1].ShapeDistTraveled
					nextShapeDistTraveled = math.Round(nextShapeDistTraveled*10) / 10

					shapePtSeqCnt++
					shapesDetail = append(shapesDetail, model.ShapeDetail{
						ShapeDetail: gtfsschedule.ShapeDetail{
							ShapeId:               id,
							ShapePtLat:            nextLat,
							ShapePtLon:            nextLon,
							ShapeDetailPtSequence: shapePtSeqCnt,
							ShapeDistTraveled:     nextShapeDistTraveled,
						},
					})
					prevShapePtLat = nextLat
					prevShapePtLon = nextLon
				}

				nextShapeDistTraveled := shapesDetail[len(shapesDetail)-1].ShapeDistTraveled + 5
				nextShapeDistTraveled = math.Round(nextShapeDistTraveled*10) / 10
				shapePtSeqCnt++
				shapesDetail = append(shapesDetail, model.ShapeDetail{
					ShapeDetail: gtfsschedule.ShapeDetail{
						ShapeId:               id,
						ShapePtLat:            nextShapePtLat,
						ShapePtLon:            nextShapePtLon,
						ShapeDetailPtSequence: shapePtSeqCnt,
						ShapeDistTraveled:     nextShapeDistTraveled,
					},
				})
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

func (g gtfsScheduleDbUseCase) createShapesDetailGeom() error {
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
				ShapeDetailGeom: gtfsschedule.ShapeDetailGeom{
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
						ShapeDetailGeom: gtfsschedule.ShapeDetailGeom{
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
						ShapeDetailGeom: gtfsschedule.ShapeDetailGeom{
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

func (g gtfsScheduleDbUseCase) createShapeEx() error {
	slog.Info("テーブル[shapes_ex] 作成開始")
	// stops, stop_timesを結合して作成したshape_exの取得（stop_idの情報はない）
	tmpShapesEx, err := g.shapeExRepo.FindShapesExByTripsAndShapes()
	if err != nil {
		return err
	}

	if err := g.shapeExRepo.CreateShapesEx(tmpShapesEx); err != nil {
		return err
	}

	tripIds, err := g.tripRepo.FindTripIds()
	if err != nil {
		return err
	}

	for _, tripId := range tripIds {
		var insertShapesEx []model.ShapeEx
		shapesEx, err := g.shapeExRepo.FindShapesExByTripId(tripId)
		if err != nil {
			return err
		}
		if len(shapesEx) == 0 {
			slog.Warn(fmt.Sprintf("trip_id: %sのshapesが存在しないためスキップします", tripId))
			continue
		}

		stopLocation, err := g.shapeExRepo.FindTripWithStopLocationByTripId(tripId)
		if err != nil {
			return err
		}

		shapesLen := len(shapesEx)
		idx := 0

		for _, stop := range stopLocation {
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

			insertShapesEx = append(insertShapesEx, model.ShapeEx{
				ShapeEx: gtfsschedule.ShapeEx{
					TripId:          tripId,
					ShapeId:         shapesEx[idx].ShapeId,
					ShapePtSequence: shapesEx[idx].ShapePtSequence,
					StopId: sql.NullString{
						String: stop.StopId,
						Valid:  true,
					},
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

func (g gtfsScheduleDbUseCase) createShapeExGeom() error {
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
				ShapeExGeom: gtfsschedule.ShapeExGeom{
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

func NewGtfsScheduleDbUseCase(
	fileManagerRepository repository.FileManagerRepository,
	gtfsScheduleRepository repository.GtfsScheduleRepository,
	gtfsScheduleGeomRepository repository.GtfsScheduleGeomRepository,
	tripRepository repository.TripRepository,
	tripGeomRepository repository.TripGeomRepository,
	shapeRepository repository.ShapeRepository,
	shapeGeomRepository repository.ShapeGeomRepository,
	shapeExRepository repository.ShapeExRepository,
	shapeExGeomRepository repository.ShapeExGeomRepository,
	shapeDetailRepository repository.ShapeDetailRepository,
	shapeDetailGeomRepository repository.ShapeDetailGeomRepository,
	stopTimeRepository repository.StopTimeRepository) GtfsScheduleDbUseCase {
	return gtfsScheduleDbUseCase{
		fileManagerRepo:      fileManagerRepository,
		gtfsScheduleRepo:     gtfsScheduleRepository,
		gtfsScheduleGeomRepo: gtfsScheduleGeomRepository,
		tripRepo:             tripRepository,
		tripGeomRepo:         tripGeomRepository,
		shapeRepo:            shapeRepository,
		shapeGeomRepo:        shapeGeomRepository,
		shapeExRepo:          shapeExRepository,
		shapeExGeomRepo:      shapeExGeomRepository,
		shapeDetailRepo:      shapeDetailRepository,
		shapeDetailGeomRepo:  shapeDetailGeomRepository,
		stopTimeRepo:         stopTimeRepository,
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
