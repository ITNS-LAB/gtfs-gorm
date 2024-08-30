package interfaces

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/infrastructure"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/usecase"
)

func GtfsDbFile(options usecase.CmdOptions) error {
	fileMangerRepository := infrastructure.NewFileManagerRepository()
	gtfsScheduleRepository := infrastructure.NewGtfsStaticRepository(options.Dsn)
	gtfsDbUseCase := usecase.NewGtfsDbUseCase(fileMangerRepository, gtfsScheduleRepository)
	if err := gtfsDbUseCase.GtfsDbFile(options); err != nil {
		return err
	}
	return nil
}

func GtfsDbUrl(options usecase.CmdOptions) error {
	fileMangerRepository := infrastructure.NewFileManagerRepository()
	gtfsScheduleRepository := infrastructure.NewGtfsStaticRepository(options.Dsn)
	gtfsDbUseCase := usecase.NewGtfsDbUseCase(fileMangerRepository, gtfsScheduleRepository)
	if err := gtfsDbUseCase.GtfsDbUrl(options); err != nil {
		return err
	}
	return nil
}
