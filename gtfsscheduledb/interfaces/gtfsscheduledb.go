package interfaces

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsscheduledb/infrastructure"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsscheduledb/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GtfsDbFile(options usecase.CmdOptions) error {
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
	//gtfsJpGeomRepository := infrastructure.NewGtfsJpGeomRepository(db)
	tripRepository := infrastructure.NewTripRepository(db)
	//tripGeomRepository := infrastructure.NewTripGeomRepository(db)
	shapeRepository := infrastructure.NewShapeRepository(db)
	//shapeGeomRepository := infrastructure.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure.NewShapeExRepository(db)
	//shapeExGeomRepository := infrastructure.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	//shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure.NewStopTimesRepository(db)

	gtfsScheduleDBuseCase := usecase.NewGtfsScheduleDbUseCase(
		fileManagerRepository,
		gtfsJpRepository,
		//gtfsJpGeomRepository,
		tripRepository,
		//tripGeomRepository,
		shapeRepository,
		//shapeGeomRepository,
		shapeExRepository,
		//shapeExGeomRepository,
		shapeDetailRepository,
		//shapeDetailGeomRepository,
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
	//gtfsJpGeomRepository := infrastructure.NewGtfsJpGeomRepository(db)
	tripRepository := infrastructure.NewTripRepository(db)
	//tripGeomRepository := infrastructure.NewTripGeomRepository(db)
	shapeRepository := infrastructure.NewShapeRepository(db)
	//shapeGeomRepository := infrastructure.NewShapeGeomRepository(db)
	shapeExRepository := infrastructure.NewShapeExRepository(db)
	//shapeExGeomRepository := infrastructure.NewShapeExGeomRepository(db)
	shapeDetailRepository := infrastructure.NewShapeDetailRepository(db)
	//shapeDetailGeomRepository := infrastructure.NewShapeDetailGeomRepository(db)
	stopTimeRepository := infrastructure.NewStopTimesRepository(db)

	gtfsScheduleDBuseCase := usecase.NewGtfsScheduleDbUseCase(
		fileManagerRepository,
		gtfsJpRepository,
		//gtfsJpGeomRepository,
		tripRepository,
		//tripGeomRepository,
		shapeRepository,
		//shapeGeomRepository,
		shapeExRepository,
		//shapeExGeomRepository,
		shapeDetailRepository,
		//shapeDetailGeomRepository,
		stopTimeRepository,
	)

	if _, err := gtfsScheduleDBuseCase.GtfsDbUrl(options); err != nil {
		return err
	}

	return nil
}
