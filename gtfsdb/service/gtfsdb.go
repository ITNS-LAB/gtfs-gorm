package service

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/infrastructure"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/usecase"
)

func GtfsDbFile(dsn, file, schema string) error {
	fileMangerRepository := infrastructure.NewFileManagerRepository()
	gtfsScheduleRepository := infrastructure.NewGtfsStaticRepository(dsn)
	gtfsDbUseCase := usecase.NewGtfsDbUseCase(fileMangerRepository, gtfsScheduleRepository)
	if err := gtfsDbUseCase.GtfsDbFile(file, schema); err != nil {
		return err
	}
	return nil
}

func GtfsDbUrl(dsn, url, schema string) error {
	fileMangerRepository := infrastructure.NewFileManagerRepository()
	gtfsScheduleRepository := infrastructure.NewGtfsStaticRepository(dsn)
	gtfsDbUseCase := usecase.NewGtfsDbUseCase(fileMangerRepository, gtfsScheduleRepository)
	if err := gtfsDbUseCase.GtfsDbUrl(url, schema); err != nil {
		return err
	}
	return nil
}

func UpdateShapes(dsn, schema string) error {
	fileMangerRepository := infrastructure.NewFileManagerRepository()
	gtfsScheduleRepository := infrastructure.NewGtfsStaticRepository(dsn)
	gtfsDbUseCase := usecase.NewGtfsDbUseCase(fileMangerRepository, gtfsScheduleRepository)
	if err := gtfsDbUseCase.RecalculateShapesUpdate(schema); err != nil {
		return err
	}
	return nil
}
