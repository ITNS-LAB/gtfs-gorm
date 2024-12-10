package usecase

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"os"
	"path"
)

type GtfsDbUseCase interface {
	GtfsDbUrl(options CmdOptions) (digest string, err error)
	GtfsDbFile(options CmdOptions) (digest string, err error)
	recalculateShapes() ([]gtfsjp.Shape, error)
	recalculateShapesUpdate(options CmdOptions) error
	createShapeEx(options CmdOptions) error
	createShapeDetail() error
	tripsGeomUpdate() error
}

type gtfsDbUseCase struct {
	fileManagerRepo  repository.FileManagerRepository
	gtfsScheduleRepo repository.GtfsScheduleRepository
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

func (g gtfsDbUseCase) GtfsDbUrl(options CmdOptions) (digest string, err error) {
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

func (g gtfsDbUseCase) GtfsDbFile(options CmdOptions) (digest string, err error) {
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
	// DB接続
	if err = g.gtfsScheduleRepo.ConnectDatabase(); err != nil {
		return "", err
	}
	// DB切断
	defer func(gtfsScheduleRepo repository.GtfsScheduleRepository) {
		_ = gtfsScheduleRepo.DisConnectDatabase()
	}(g.gtfsScheduleRepo)
	// スキーマの移動
	if err = g.gtfsScheduleRepo.SetSchema(options.Schema); err != nil {
		return "", err
	}
	// マイグレーション
	if err = g.gtfsScheduleRepo.Migrate(options.ShapesEx, options.ShapesDetail); err != nil {
		return "", err
	}
	// データ挿入
	if err = g.gtfsScheduleRepo.Create(gtfsPath); err != nil {
		return "", err
	}

	//// tripsにgeomデータを追加
	//if err = g.tripsGeomUpdate(); err != nil {
	//	return "", err
	//}
	//
	//// optionsの実行
	//if options.RecalculateDist {
	//	if err = g.recalculateShapesUpdate(options); err != nil {
	//		return "", err
	//	}
	//}
	//
	//if options.Geom {
	//	if err = g.tripsGeomUpdate(); err != nil {
	//		return "", err
	//	}
	//}
	//if options.ShapesEx {
	//	if err = g.createShapeEx(options); err != nil {
	//		return digest, err
	//	}
	//}
	//if options.ShapesDetail {
	//	if err = g.createShapeDetail(); err != nil {
	//		return digest, err
	//	}
	//}
	return digest, err
}

//func (g gtfsDbUseCase) recalculateShapes() ([]gtfsjp.Shape, error) {
//	var res []gtfsjp.Shape
//
//	slog.Info("テーブル[shapes] shape_dist_traveled の再計算を行います。")
//
//	// shape_idのスライスを取得
//	shapeIds, err := g.gtfsScheduleRepo.FindShapeIds()
//	if err != nil {
//		return nil, err
//	}
//
//	// shape_dist_traveledの追加
//	for _, shapeId := range shapeIds {
//		totalDistance := 0.0
//		shapes, err := g.gtfsScheduleRepo.FindShapes(shapeId)
//		if err != nil {
//			return nil, err
//		}
//		for i, pt := range shapes {
//			if i == 0 {
//				distTraveled := math.Floor(totalDistance)
//				pt.ShapeDistTraveled = &distTraveled
//				res = append(res, pt)
//				continue
//			}
//			// 2点間の距離を計算
//			totalDistance += util.KarneyWgs84(*shapes[i-1].ShapePtLat, *shapes[i-1].ShapePtLon, *pt.ShapePtLat, *pt.ShapePtLon)
//			distTraveled := math.Floor(totalDistance)
//			pt.ShapeDistTraveled = &distTraveled
//			res = append(res, pt)
//		}
//	}
//	slog.Info("テーブル[shapes] shape_dist_traveled の再計算が完了しました。")
//	return res, nil
//}
//
//func (g gtfsDbUseCase) recalculateShapesUpdate(options CmdOptions) error {
//	shapes, err := g.recalculateShapes()
//	if err != nil {
//		return err
//	}
//
//	slog.Info("テーブル[shapes] shape_dist_traveled の更新を行います。")
//
//	if err := g.gtfsScheduleRepo.UpdateShapes(shapes); err != nil {
//		return err
//	}
//	slog.Info("テーブル[shapes] shape_dist_traveled の更新が完了しました。")
//	return nil
//}
//
//// shapeExとstop_timesの更新が1つのメソッドで行われているので分割したい
//func (g gtfsDbUseCase) createShapeEx(options CmdOptions) error {
//	tmpShapesEx, err := g.gtfsScheduleRepo.FetchShapesWithTrips()
//	if err != nil {
//		return err
//	}
//
//	if err := g.gtfsScheduleRepo.CreateShapesEx(tmpShapesEx); err != nil {
//		return err
//	}
//	slog.Info("shapes_exテーブルを作成しました。")
//
//	tripIds, err := g.gtfsScheduleRepo.FindTripIds()
//	if err != nil {
//		return err
//	}
//
//	slog.Info("stop_timesテーブルshape_dist_traveledの更新を開始しました。")
//	slog.Info("shape_exテーブルstop_idの更新を開始しました。")
//	for _, tripId := range tripIds {
//		var shapesEx []gtfsjp.ShapeEx
//		var stopTimes []gtfsjp.StopTime
//
//		shapes, err := g.gtfsScheduleRepo.FindShapesWithTripsByTripId(tripId)
//		if err != nil {
//			return err
//		}
//		stopTimesWithLocations, err := g.gtfsScheduleRepo.FindStopTimesByTripId(tripId)
//		if err != nil {
//			return err
//		}
//
//		shapesLen := len(shapes)
//		tmpIdx := 0
//		for _, stop := range stopTimesWithLocations {
//			minDist := math.MaxFloat64
//			for i := tmpIdx; i < shapesLen; i++ {
//				dist := util.KarneyWgs84(*stop.StopLat, *stop.StopLon, *shapes[i].ShapePtLat, *shapes[i].ShapePtLon)
//				if dist < minDist {
//					minDist = dist
//					tmpIdx = i
//				}
//				if minDist <= 0 {
//					break
//				}
//			}
//
//			stopTimes = append(stopTimes, gtfsjp.StopTime{
//				TripId:            shapes[tmpIdx].TripId,
//				StopId:            stop.StopId,
//				StopSequence:      stop.StopSequence,
//				ShapeDistTraveled: shapes[tmpIdx].ShapeDistTraveled,
//			})
//
//			shapesEx = append(shapesEx, gtfsjp.ShapeEx{
//				TripId:          shapes[tmpIdx].TripId,
//				ShapeId:         shapes[tmpIdx].ShapeId,
//				ShapePtSequence: shapes[tmpIdx].ShapePtSequence,
//				StopId:          stop.StopId,
//			})
//		}
//
//		if err := g.gtfsScheduleRepo.UpdateShapesEx(shapesEx); err != nil {
//			return err
//		}
//
//		if err := g.gtfsScheduleRepo.UpdateStopTimes(stopTimes); err != nil {
//			return err
//		}
//	}
//	slog.Info("stop_timesテーブルshape_dist_traveledの更新が完了しました。")
//	slog.Info("shape_exテーブルstop_idの更新が完了しました。")
//	return nil
//}
//
//func (g gtfsDbUseCase) tripsGeomUpdate() error {
//	slog.Info("テーブル[trips] tripsテーブルのgeomを更新します。")
//	// shape_idの取得
//	shapeIds, err := g.gtfsScheduleRepo.FindShapeIds()
//	if err != nil {
//		return err
//	}
//	for _, shapeId := range shapeIds {
//		// 特定のshape_idのshapesを取得
//		shapes, err := g.gtfsScheduleRepo.FindShapes(shapeId)
//		if err != nil {
//			return err
//		}
//
//		// 特定のshape_idのgeom(LineString)を作成
//		var lineString orb.LineString
//		for _, pt := range shapes {
//			lineString = append(lineString, orb.Point{*pt.ShapePtLon, *pt.ShapePtLat})
//		}
//
//		// 特定のshape_idのtripsを取得
//		trips, err := g.gtfsScheduleRepo.FindTripsByShapeId(shapeId)
//		if err != nil {
//			return err
//		}
//		for i := range trips {
//			geom := gormdatatypes.Geometry{
//				Geom: lineString,
//				Srid: 4326,
//			}
//			trips[i].Geom = &geom
//		}
//		if err := g.gtfsScheduleRepo.UpdateTrips(trips); err != nil {
//			return err
//		}
//	}
//	slog.Info("テーブル[trips] tripsテーブルのgeomを更新が完了しました。")
//	return nil
//}
//
//func (g gtfsDbUseCase) createShapeDetail() error {
//	shapeIds, err := g.gtfsScheduleRepo.FindShapeIds()
//	if err != nil {
//		return err
//	}
//
//	interval := 5.0
//
//	for _, shapeId := range shapeIds {
//		shapes, err := g.gtfsScheduleRepo.FindShapes(shapeId)
//		if err != nil {
//			return err
//		}
//
//		shapesDetail, err := g.resampleShapeDetail(shapes, interval)
//		if err != nil {
//			return err
//		}
//
//		if err := g.gtfsScheduleRepo.CreateShapeDetail(shapesDetail); err != nil {
//			return err
//		}
//
//	}
//
//	return nil
//}
//
//func (g gtfsDbUseCase) resampleShapeDetail(shapes []gtfsjp.Shape, interval float64) ([]gtfsjp.ShapeDetail, error) {
//	shapesDetail := []gtfsjp.ShapeDetail{{
//		ShapeId:               shapes[0].ShapeId,
//		ShapePtLat:            shapes[0].ShapePtLat,
//		ShapePtLon:            shapes[0].ShapePtLon,
//		ShapeDetailPtSequence: shapes[0].ShapePtSequence,
//		ShapeDistTraveled:     shapes[0].ShapeDistTraveled,
//		Geom:                  shapes[0].Geom}}
//
//	// 残り
//	remainder := 0.0
//	shapePtSequenceCounter := *shapes[0].ShapePtSequence
//
//	for i := 1; i < len(shapes); i++ {
//		// 1つ前のshape_pt
//		prevShapePtLat := *shapesDetail[len(shapesDetail)-1].ShapePtLat
//		prevShapePtLon := *shapesDetail[len(shapesDetail)-1].ShapePtLon
//		// 1つ次のshape_pt
//		nextShapePtLat := *shapes[i].ShapePtLat
//		nextShapePtLon := *shapes[i].ShapePtLon
//
//		// prevShapePtとnextShapePtの距離(区間距離)
//		blockDistance := util.KarneyWgs84(prevShapePtLat, prevShapePtLon, nextShapePtLat, nextShapePtLon)
//		blockDistance += remainder
//
//		// 区間距離からshapesを分割する回数を計算
//		repeat := int(blockDistance / interval)
//		remainder = math.Mod(blockDistance, interval)
//
//		// t:媒介変数 dLat,dLon:方向ベクトル
//		t := interval / blockDistance
//		dLat := nextShapePtLat - prevShapePtLat
//		dLon := nextShapePtLon - prevShapePtLon
//
//		for j := 0; j < repeat; j++ {
//			shapePtSequenceCounter++
//			prevLat := *shapesDetail[len(shapesDetail)-1].ShapePtLat
//			prevLon := *shapesDetail[len(shapesDetail)-1].ShapePtLon
//			nextLat := t*dLat + prevLat
//			nextLon := t*dLon + prevLon
//
//			shortDistance := util.KarneyWgs84(prevLat, prevLon, nextLat, nextLon)
//
//			shapeDistTraveled := *shapesDetail[len(shapesDetail)-1].ShapeDistTraveled + shortDistance
//			geomPoint := orb.Point{nextLat, nextLon}
//			shapesDetail = append(shapesDetail, gtfsjp.ShapeDetail{
//				ShapeId:               shapes[0].ShapeId,
//				ShapePtLat:            ptr.Ptr(nextLat),
//				ShapePtLon:            ptr.Ptr(nextLon),
//				ShapeDetailPtSequence: ptr.Ptr(shapePtSequenceCounter),
//				ShapeDistTraveled:     ptr.Ptr(math.Round(shapeDistTraveled*100) / 100),
//				Geom:                  &geomdatatypes.Geometry{Geom: geomPoint, Srid: 4326},
//			})
//
//		}
//
//	}
//	return shapesDetail, nil
//}

func (g gtfsDbUseCase) recalculateShapes() ([]gtfsjp.Shape, error) {
	//TODO implement me
	panic("implement me")
}

func (g gtfsDbUseCase) recalculateShapesUpdate(options CmdOptions) error {
	//TODO implement me
	panic("implement me")
}

func (g gtfsDbUseCase) createShapeEx(options CmdOptions) error {
	//TODO implement me
	panic("implement me")
}

func (g gtfsDbUseCase) createShapeDetail() error {
	//TODO implement me
	panic("implement me")
}

func (g gtfsDbUseCase) tripsGeomUpdate() error {
	//TODO implement me
	panic("implement me")
}

func NewGtfsDbUseCase(fileMangerRepo repository.FileManagerRepository, gtfsScheduleRepo repository.GtfsScheduleRepository) GtfsDbUseCase {
	return &gtfsDbUseCase{
		fileManagerRepo:  fileMangerRepo,
		gtfsScheduleRepo: gtfsScheduleRepo,
	}
}
