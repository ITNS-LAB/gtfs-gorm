package repository

import "github.com/ITNS-LAB/gtfs-gorm/ormstatic"

type GtfsScheduleRepository interface {
	ConnectDatabase() error
	DisConnectDatabase() error
	Migrate() error
	Create(gtfsPath string) error
	CreateSchema(schema string) error
	SetSchema(schema string) error
	FindShapeIds() ([]string, error)
	FindShapes(shapeId string) ([]ormstatic.Shape, error)
	UpdateShapes([]ormstatic.Shape) error
}
