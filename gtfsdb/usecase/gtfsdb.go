package usecase

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/model"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"log/slog"
	"math"
	"os"
	"path"
)

type GtfsJpDbUseCase interface {
	GtfsDbUrl(options CmdOptions) (digest string, err error)
	GtfsDbFile(options CmdOptions) (digest string, err error)
	recalculateShapes() error
	recalculateShapesGeom() error
}

type gtfsJpDbUseCase struct {
	fileManagerRepo repository.FileManagerRepository
	gtfsJpRepo      repository.GtfsJpRepository
	gtfsJpGeomRepo  repository.GtfsJpGeomRepository
	tripRepo        repository.TripRepository
	tripGeomRepo    repository.TripGeomRepository
	shapeRepo       repository.ShapeRepository
	shapeGeomRepo   repository.ShapeGeomRepository
	shapeExRepo     repository.ShapeExRepository
	shapeDetailRepo repository.ShapeDetailRepository
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
		if options.ShapesEx {
			if err = g.shapeExRepo.MigrateShapesEx(); err != nil {
				return "", err
			}
			// ビジネスロジック
			// shapesの取得
			//shapes, err := g.shapeGeomRepo.FetchShapesGeom()
			//if err != nil {
			//	return "", err
			//}
			//fmt.Println(shapes)
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

	}

	if options.ShapesDetail {

	}

	return digest, err
}

func (g gtfsJpDbUseCase) recalculateShapes() error {
	slog.Info("テーブル[shapes] shape_dist_traveled の再計算を実行します")

	// shape_idのスライスを取得
	shapeIds, err := g.shapeRepo.FindShapeIds()
	if err != nil {
		return err
	}

	for _, shapeId := range shapeIds {
		var recalculateShapes []model.Shape
		totalDistance := 0.0
		shapes, err := g.shapeRepo.FindShapes(shapeId)
		if err != nil {
			return err
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
			totalDistance += util.KarneyWgs84(shapes[i-1].ShapePtLat, shapes[i-1].ShapePtLon, pt.ShapePtLat, pt.ShapePtLon)
			distTraveled := math.Floor(totalDistance)
			pt.ShapeDistTraveled = &distTraveled
			recalculateShapes = append(recalculateShapes, pt)
		}

		// 再計算したデータをDBに格納
		if err := g.shapeRepo.UpdateShapes(recalculateShapes); err != nil {
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

	for _, shapeId := range shapeIds {
		var recalculateShapes []model.ShapeGeom
		totalDistance := 0.0
		shapes, err := g.shapeGeomRepo.FindShapesGeom(shapeId)
		if err != nil {
			return err
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
			totalDistance += util.KarneyWgs84(shapes[i-1].ShapePtLat, shapes[i-1].ShapePtLon, pt.ShapePtLat, pt.ShapePtLon)
			distTraveled := math.Floor(totalDistance)
			pt.ShapeDistTraveled = &distTraveled
			recalculateShapes = append(recalculateShapes, pt)
		}

		// 再計算したデータをDBに格納
		if err := g.shapeGeomRepo.UpdateShapesGeom(recalculateShapes); err != nil {
			return err
		}
	}
	slog.Info("テーブル[shapes] shape_dist_traveled の再計算完了")
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
	shapeDetailRepository repository.ShapeDetailRepository) GtfsJpDbUseCase {
	return gtfsJpDbUseCase{
		fileManagerRepo: fileManagerRepository,
		gtfsJpRepo:      gtfsJpRepository,
		gtfsJpGeomRepo:  gtfsJpGeomRepository,
		tripRepo:        tripRepository,
		tripGeomRepo:    tripGeomRepository,
		shapeRepo:       shapeRepository,
		shapeGeomRepo:   shapeGeomRepository,
		shapeExRepo:     shapeExRepository,
		shapeDetailRepo: shapeDetailRepository,
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
