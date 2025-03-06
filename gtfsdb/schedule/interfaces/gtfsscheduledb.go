package interfaces

import (
	"fmt"
	infrastructure2 "github.com/ITNS-LAB/gtfs-gorm/gtfsdb/schedule/infrastructure"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/schedule/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GtfsDbFile(options usecase.CmdOptions) error {
	// db接続	"%s?search_path=%s"の?を&に変更している
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s&search_path=%s", options.Dsn, options.Schema)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	// DI
	fileManagerRepository := infrastructure2.NewFileManagerRepository()
	gtfsScheduleRepository := infrastructure2.NewGtfsScheduleRepository(db)
	gtfsScheduleGeomRepository := infrastructure2.NewGtfsScheduleGeomRepository(db)
	tripRepository := infrastructure2.NewTripRepository(db)
	tripGeomRepository := infrastructure2.NewTripGeomRepository(db)
	shapeRepository := infrastructure2.NewShapeRepository(db)
	shapeGeomRepository := infrastructure2.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure2.NewShapeExRepository(db)
	shapeExGeomRepository := infrastructure2.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure2.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure2.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure2.NewStopTimesRepository(db)

	gtfsScheduleDBuseCase := usecase.NewGtfsScheduleDbUseCase(
		fileManagerRepository,
		gtfsScheduleRepository,
		gtfsScheduleGeomRepository,
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

	if _, err := gtfsScheduleDBuseCase.GtfsDbFile(options); err != nil {
		return err
	}

	return nil
}

func GtfsDbUrl(options usecase.CmdOptions) error {
	// db接続
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s?search_path=%s", options.Dsn, options.Schema)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	// DI注入
	fileManagerRepository := infrastructure2.NewFileManagerRepository()
	gtfsJpRepository := infrastructure2.NewGtfsScheduleRepository(db)
	gtfsScheduleGeomRepository := infrastructure2.NewGtfsScheduleGeomRepository(db)
	tripRepository := infrastructure2.NewTripRepository(db)
	tripGeomRepository := infrastructure2.NewTripGeomRepository(db)
	shapeRepository := infrastructure2.NewShapeRepository(db)
	shapeGeomRepository := infrastructure2.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure2.NewShapeExRepository(db)
	shapeExGeomRepository := infrastructure2.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure2.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure2.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure2.NewStopTimesRepository(db)
	gtfsScheduleDBuseCase := usecase.NewGtfsScheduleDbUseCase(
		fileManagerRepository,
		gtfsJpRepository,
		gtfsScheduleGeomRepository,
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

	if _, err := gtfsScheduleDBuseCase.GtfsDbUrl(options); err != nil {
		return err
	}

	return nil
}
