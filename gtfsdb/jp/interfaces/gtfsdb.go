package interfaces

import (
	"errors"
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/jp/infrastructure"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/jp/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GtfsDbFile(options usecase.CmdOptions) error {
	a := 1
	if a == 1 {
		return errors.New("db接続より前まで実行できているよ！！(file)")
	}

	// db接続
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s?search_path=%s", options.Dsn, options.Schema)), &gorm.Config{
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
	shapeExGeomRepository := infrastructure.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure.NewStopTimesRepository(db)
	gtfsJpDBuseCase := usecase.NewGtfsJpDbUseCase(
		fileManagerRepository,
		gtfsJpRepository,
		gtfsJpGeomRepository,
		tripRepository,
		tripGeomRepository,
		shapeRepository,
		shapeGeomRepository,
		shapeExRepository,
		shapeExGeomRepository,
		shapeDetailRepository,
		shapeDetailGeomRepository,
		stopTimeRepository,
	)

	if _, err := gtfsJpDBuseCase.GtfsDbFile(options); err != nil {
		return err
	}

	return nil
}

func GtfsDbUrl(options usecase.CmdOptions) error {
	a := 1
	if a == 1 {
		return errors.New("db接続より前まで実行できているよ！！(url)")
	}
	// db接続
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s?search_path=%s", options.Dsn, options.Schema)), &gorm.Config{
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
	shapeExGeomRepository := infrastructure.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure.NewStopTimesRepository(db)
	gtfsJpDBuseCase := usecase.NewGtfsJpDbUseCase(
		fileManagerRepository,
		gtfsJpRepository,
		gtfsJpGeomRepository,
		tripRepository,
		tripGeomRepository,
		shapeRepository,
		shapeGeomRepository,
		shapeExRepository,
		shapeExGeomRepository,
		shapeDetailRepository,
		shapeDetailGeomRepository,
		stopTimeRepository,
	)

	if _, err := gtfsJpDBuseCase.GtfsDbUrl(options); err != nil {
		return err
	}

	return nil
}
