package interfaces

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/infrastructure"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GtfsDbFile(options usecase.CmdOptions) error {
	// db接続
	db, err := gorm.Open(postgres.Open(options.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	// DI注入
	fileManagerRepository := infrastructure.NewFileManagerRepository()
	gtfsJpRepository := infrastructure.NewGtfsJpRepository(db)
	gtfsJpGeomRepository := infrastructure.NewGtfsJpGeomRepository(db)
	tripRepository := infrastructure.NewTripRepository(db)
	tripGeomRepository := infrastructure.NewTripGeomRepository(db)
	shapeRepository := infrastructure.NewShapeRepository(db)
	shapeGeomRepository := infrastructure.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure.NewShapeExRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	gtfsJpDBuseCase := usecase.NewGtfsJpDbUseCase(
		fileManagerRepository,
		gtfsJpRepository,
		gtfsJpGeomRepository,
		tripRepository,
		tripGeomRepository,
		shapeRepository,
		shapeGeomRepository,
		shapeExRepository,
		shapeDetailRepository,
		shapeDetailGeomRepository,
	)

	if _, err := gtfsJpDBuseCase.GtfsDbFile(options); err != nil {
		return err
	}

	return nil
}

func GtfsDbUrl(options usecase.CmdOptions) error {
	// db接続
	db, err := gorm.Open(postgres.Open(options.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	// DI注入
	fileManagerRepository := infrastructure.NewFileManagerRepository()
	gtfsJpRepository := infrastructure.NewGtfsJpRepository(db)
	gtfsJpGeomRepository := infrastructure.NewGtfsJpGeomRepository(db)
	tripRepository := infrastructure.NewTripRepository(db)
	tripGeomRepository := infrastructure.NewTripGeomRepository(db)
	shapeRepository := infrastructure.NewShapeRepository(db)
	shapeGeomRepository := infrastructure.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure.NewShapeExRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	gtfsJpDBuseCase := usecase.NewGtfsJpDbUseCase(
		fileManagerRepository,
		gtfsJpRepository,
		gtfsJpGeomRepository,
		tripRepository,
		tripGeomRepository,
		shapeRepository,
		shapeGeomRepository,
		shapeExRepository,
		shapeDetailRepository,
		shapeDetailGeomRepository,
	)

	if _, err := gtfsJpDBuseCase.GtfsDbUrl(options); err != nil {
		return err
	}

	return nil
}
