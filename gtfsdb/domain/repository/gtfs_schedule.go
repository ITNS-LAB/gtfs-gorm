package repository

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/model"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
)

type GtfsScheduleRepository interface {
	ConnectDatabase() error
	DisConnectDatabase() error
	Migrate(shapeEx bool, shapeDetail bool) error
	Create(gtfsPath string) error
	CreateSchema(schema string) error
	SetSchema(schema string) error
	FindShapeIds() ([]string, error)
	FindShapes(shapeId string) ([]ormstatic.Shape, error)
	UpdateShapes([]ormstatic.Shape) error
	FindTripIds() ([]string, error)
	FindShapesWithTripsByTripId(tripId string) ([]ormstatic.ShapeEx, error)
	FindStopTimesByTripId(tripId string) ([]model.StopTimeWithLocation, error)
	CreateShapesEx([]ormstatic.ShapeEx) error
	UpdateShapesEx([]ormstatic.ShapeEx) error
	FetchShapes() ([]ormstatic.Shape, error)
	FetchShapesWithTrips() ([]ormstatic.ShapeEx, error)
	UpdateStopTimes([]ormstatic.StopTime) error
	FindTripsByShapeId(shapeId string) ([]ormstatic.Trip, error)
	UpdateTrips([]ormstatic.Trip) error
	CreateShapeDetail([]ormstatic.ShapeDetail) error
}
