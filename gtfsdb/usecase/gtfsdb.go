package usecase

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"gorm.io/gorm"
	"os"
	"path"
)

type GtfsJpDbUseCase interface {
	GtfsDbUrl(db *gorm.DB, options CmdOptions) (digest string, err error)
	GtfsDbFile(db *gorm.DB, options CmdOptions) (digest string, err error)
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

func (g gtfsJpDbUseCase) GtfsDbUrl(db *gorm.DB, options CmdOptions) (digest string, err error) {
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
	if digest, err = g.GtfsDbFile(db, options); err != nil {
		return "", err
	}
	return digest, nil
}

func (g gtfsJpDbUseCase) GtfsDbFile(db *gorm.DB, options CmdOptions) (digest string, err error) {
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
		if err = g.gtfsJpGeomRepo.Migrate(); err != nil {
			return "", err
		}
		if err = g.gtfsJpGeomRepo.Create(gtfsPath); err != nil {
			return "", err
		}
		//if options.ShapesEx {
		//	if err = g.shapeExRepo.Migrate(); err != nil {
		//		return "", err
		//	}
		//
		//}
	} else {
		if err = g.gtfsJpRepo.Migrate(); err != nil {
			return "", err
		}
		if err = g.gtfsJpRepo.Create(gtfsPath); err != nil {
			return "", err
		}
	}

	if options.ShapesDetail {

	}

	return digest, err
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
