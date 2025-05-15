package interfaces

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/schedule/infrastructure"
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
	fileManagerRepository := infrastructure.NewFileManagerRepository()
	gtfsScheduleRepository := infrastructure.NewGtfsScheduleRepository(db)
	gtfsScheduleGeomRepository := infrastructure.NewGtfsScheduleGeomRepository(db)
	tripRepository := infrastructure.NewTripRepository(db)
	tripGeomRepository := infrastructure.NewTripGeomRepository(db)
	shapeRepository := infrastructure.NewShapeRepository(db)
	shapeGeomRepository := infrastructure.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure.NewShapeExRepository(db)
	shapeExGeomRepository := infrastructure.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure.NewStopTimesRepository(db)

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
	fileManagerRepository := infrastructure.NewFileManagerRepository()
	gtfsJpRepository := infrastructure.NewGtfsScheduleRepository(db)
	gtfsScheduleGeomRepository := infrastructure.NewGtfsScheduleGeomRepository(db)
	tripRepository := infrastructure.NewTripRepository(db)
	tripGeomRepository := infrastructure.NewTripGeomRepository(db)
	shapeRepository := infrastructure.NewShapeRepository(db)
	shapeGeomRepository := infrastructure.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure.NewShapeExRepository(db)
	shapeExGeomRepository := infrastructure.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure.NewStopTimesRepository(db)
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
