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
	GtfsDbUrl(options CmdOptions) (digest string, err error)
	GtfsDbFile(options CmdOptions) (digest string, err error)
	recalculateShapes(options CmdOptions) ([]ormstatic.Shape, error)
	recalculateShapesUpdate(options CmdOptions) error
}

type gtfsDbUseCase struct {
	fileManagerRepo  repository.FileManagerRepository
	gtfsScheduleRepo repository.GtfsScheduleRepository
}

type CmdOptions struct {
	GtfsUrl         string
	GtfsFile        string
	ShapesEx        bool
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
	if err = g.gtfsScheduleRepo.Migrate(); err != nil {
		return "", err
	}
	// データ挿入
	if err = g.gtfsScheduleRepo.Create(gtfsPath); err != nil {
		return "", err
	}

	// optionsの実行
	if options.RecalculateDist {
		if err = g.recalculateShapesUpdate(options); err != nil {
			return "", err
		}
	}
	return digest, err
}

func (g gtfsDbUseCase) recalculateShapes(options CmdOptions) ([]ormstatic.Shape, error) {
	var res []ormstatic.Shape

	slog.Info("テーブル[shapes] shape_dist_traveled の再計算を行います。")

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
	slog.Info("テーブル[shapes] shape_dist_traveled の再計算が完了しました。")
	return res, nil
}

func (g gtfsDbUseCase) recalculateShapesUpdate(options CmdOptions) error {
	shapes, err := g.recalculateShapes(options)
	if err != nil {
		return err
	}

	slog.Info("テーブル[shapes] shape_dist_traveled の更新を行います。")

	if err := g.gtfsScheduleRepo.UpdateShapes(shapes); err != nil {
		return err
	}
	slog.Info("テーブル[shapes] shape_dist_traveled の更新が完了しました。")
	return nil
}

func NewGtfsDbUseCase(fileMangerRepo repository.FileManagerRepository, gtfsScheduleRepo repository.GtfsScheduleRepository) GtfsDbUseCase {
	return &gtfsDbUseCase{
		fileManagerRepo:  fileMangerRepo,
		gtfsScheduleRepo: gtfsScheduleRepo,
	}
}
