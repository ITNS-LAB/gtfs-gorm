package repository

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/model"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
)

type GtfsScheduleRepository interface {
	ConnectDatabase() error
	DisConnectDatabase() error
	Migrate(shapeEx bool, shapeDetail bool) error
	Create(gtfsPath string) error
	CreateSchema(schema string) error
	SetSchema(schema string) error
	FindShapeIds() ([]string, error)
	FindShapes(shapeId string) ([]gtfsjp.Shape, error)
	UpdateShapes([]gtfsjp.Shape) error
	FindTripIds() ([]string, error)
	FindShapesWithTripsByTripId(tripId string) ([]gtfsjp.ShapeEx, error)
	FindStopTimesByTripId(tripId string) ([]model.StopTimeWithLocation, error)
	CreateShapesEx([]gtfsjp.ShapeEx) error
	UpdateShapesEx([]gtfsjp.ShapeEx) error
	FetchShapes() ([]gtfsjp.Shape, error)
	FetchShapesWithTrips() ([]gtfsjp.ShapeEx, error)
	UpdateStopTimes([]gtfsjp.StopTime) error
	FindTripsByShapeId(shapeId string) ([]gtfsjp.Trip, error)
	UpdateTrips([]gtfsjp.Trip) error
	CreateShapeDetail([]gtfsjp.ShapeDetail) error
}
