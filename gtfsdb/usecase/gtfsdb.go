package usecase

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"log/slog"
	"math"
	"os"
	"path"
)

type GtfsDbUseCase interface {
	GtfsDbUrl(url, schema string) error
	GtfsDbFile(file, schema string) error
	recalculateShapes(schema string) ([]ormstatic.Shape, error)
	RecalculateShapesUpdate(schema string) error
}

type gtfsDbUseCase struct {
	fileManagerRepo  repository.FileManagerRepository
	gtfsScheduleRepo repository.GtfsScheduleRepository
}

func (g gtfsDbUseCase) GtfsDbUrl(url, schema string) error {
	// tmpディレクトリを作成
	tmp := "tmp"
	if err := os.MkdirAll(tmp, 0755); err != nil {
		return err
	}
	// gtfsをダウンロード
	gtfsZip := path.Join(tmp, "gtfs.zip")
	if err := g.fileManagerRepo.Download(url, gtfsZip); err != nil {
		return err
	}
	if err := g.GtfsDbFile(gtfsZip, schema); err != nil {
		return err
	}
	return nil
}

func (g gtfsDbUseCase) GtfsDbFile(file, schema string) error {
	// tmpディレクトリを作成
	tmp := "tmp"
	if err := os.MkdirAll(tmp, 0755); err != nil {
		return err
	}
	// tmpディレクトリの削除
	defer func(fileManagerRepo repository.FileManagerRepository, path string) {
		_ = fileManagerRepo.Remove(path)
	}(g.fileManagerRepo, tmp)
	// gtfsを解凍
	gtfsPath, err := g.fileManagerRepo.UnZip(file, tmp)
	if err != nil {
		return err
	}
	// DB接続
	if err := g.gtfsScheduleRepo.ConnectDatabase(); err != nil {
		return err
	}
	// DB切断
	defer func(gtfsScheduleRepo repository.GtfsScheduleRepository) {
		_ = gtfsScheduleRepo.DisConnectDatabase()
	}(g.gtfsScheduleRepo)
	// スキーマの移動
	if err := g.gtfsScheduleRepo.SetSchema(schema); err != nil {
		return err
	}
	// マイグレーション
	if err := g.gtfsScheduleRepo.Migrate(); err != nil {
		return err
	}
	// データ挿入
	if err := g.gtfsScheduleRepo.Create(gtfsPath); err != nil {
		return err
	}
	return nil
}

func (g gtfsDbUseCase) recalculateShapes(schema string) ([]ormstatic.Shape, error) {
	var res []ormstatic.Shape

	slog.Info("shapes shape_dist_traveled の再計算を行います。")

	// DB接続
	if err := g.gtfsScheduleRepo.ConnectDatabase(); err != nil {
		return nil, err
	}
	// DB切断
	defer func(gtfsScheduleRepo repository.GtfsScheduleRepository) {
		_ = gtfsScheduleRepo.DisConnectDatabase()
	}(g.gtfsScheduleRepo)

	// スキーマの移動
	if err := g.gtfsScheduleRepo.SetSchema(schema); err != nil {
		return nil, err
	}

	// shape_idのスライスを取得
	shapeIds, err := g.gtfsScheduleRepo.ReadShapeIds()
	if err != nil {
		return nil, err
	}

	// shape_dist_traveledの追加
	for _, shapeId := range shapeIds {
		totalDistance := 0.0
		shapes, err := g.gtfsScheduleRepo.ReadShapes(shapeId)
		if err != nil {
			return nil, err
		}
		for i, pt := range shapes {
			if i == 0 {
				distTraveled := math.Floor(totalDistance)
				pt.ShapeDistTraveled = &distTraveled
				res = append(res, pt)
				continue
			}
			// 2点間の距離を計算
			totalDistance += util.KarneyWgs84(*shapes[i-1].ShapePtLat, *shapes[i-1].ShapePtLon, *pt.ShapePtLat, *pt.ShapePtLon)
			distTraveled := math.Floor(totalDistance)
			pt.ShapeDistTraveled = &distTraveled
			res = append(res, pt)
		}
	}
	slog.Info("shapes shape_dist_traveled の再計算が完了しました。")
	return res, nil
}

func (g gtfsDbUseCase) RecalculateShapesUpdate(schema string) error {
	shapes, err := g.recalculateShapes(schema)
	if err != nil {
		return err
	}

	slog.Info("shapes shape_dist_traveled の更新を行います。")
	// DB接続
	if err := g.gtfsScheduleRepo.ConnectDatabase(); err != nil {
		return err
	}
	// DB切断
	defer func(gtfsScheduleRepo repository.GtfsScheduleRepository) {
		_ = gtfsScheduleRepo.DisConnectDatabase()
	}(g.gtfsScheduleRepo)

	// スキーマの移動
	if err := g.gtfsScheduleRepo.SetSchema(schema); err != nil {
		return err
	}

	if err := g.gtfsScheduleRepo.UpdateShapes(shapes); err != nil {
		return err
	}
	slog.Info("shapes shape_dist_traveled の更新が完了しました。")
	return nil
}

func NewGtfsDbUseCase(fileMangerRepo repository.FileManagerRepository, gtfsScheduleRepo repository.GtfsScheduleRepository) GtfsDbUseCase {
	return &gtfsDbUseCase{
		fileManagerRepo:  fileMangerRepo,
		gtfsScheduleRepo: gtfsScheduleRepo,
	}
}
