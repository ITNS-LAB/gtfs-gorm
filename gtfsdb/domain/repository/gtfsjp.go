package repository

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/domain/model"
)

type GtfsJpRepository interface {
	Migrate() error
	Create(gtfsPath string) error
}

type GtfsJpGeomRepository interface {
	Migrate() error
	Create(gtfsPath string) error
}

type TripRepository interface {
	FindTripIds() ([]string, error)
}

type TripGeomRepository interface {
	FindTripsGeomIds() ([]string, error)
	UpdateTripsGeom([]model.TripGeom) error
}

//type StopTimeRepository interface {
//	UpdateStopTimes([]model.StopTime)
//}

type ShapeRepository interface {
	FindShapeIds() ([]string, error)
	FindShapes(shapeId string) ([]model.Shape, error)
	UpdateShapes([]model.Shape) error
	FetchShapes() ([]model.Shape, error)
}

type ShapeGeomRepository interface {
	FindShapeGeomIds() ([]string, error)
	FindShapesGeom(shapeId string) ([]model.ShapeGeom, error)
	UpdateShapesGeom([]model.ShapeGeom) error
	FetchShapesGeom() ([]model.ShapeGeom, error)
}

type ShapeExRepository interface {
	Migrate() error
	CreateShapesEx([]model.ShapeEx) error
	UpdateShapesEx([]model.ShapeEx) error
}

type ShapeDetailRepository interface {
	Migrate() error
	CreateShapeDetail([]model.ShapeDetail) error
}
