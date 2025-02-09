package repository

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsscheduledb/domain/model"
)

type GtfsScheduleRepository interface {
	MigrateGtfsSchedule() error
	CreateGtfsSchedule(gtfsPath string) error
}

/*
type GtfsJpGeomRepository interface {
	MigrateGtfsJpGeom() error
	CreateGtfsJpGeom(gtfsPath string) error
}
*/

type TripRepository interface {
	FindTripIds() ([]string, error)
	FindShapeIdByTripId(tripId string) (string, error)
}

/*
type TripGeomRepository interface {
	FindTripsGeomIds() ([]string, error)
	UpdateTripsGeom([]model.TripGeom) error
}
*/

type ShapeRepository interface {
	FindShapeIds() ([]string, error)
	FindShapesByShapeId(shapeId string) ([]model.Shape, error)
	UpdateShapes([]model.Shape) error
	FindShapes() ([]model.Shape, error)
}

/*
type ShapeGeomRepository interface {
	FindShapeGeomIds() ([]string, error)
	FindShapesGeomByShapeId(shapeId string) ([]model.ShapeGeom, error)
	UpdateShapesGeom([]model.ShapeGeom) error
	FindShapesGeom() ([]model.ShapeGeom, error)
}
*/

type ShapeExRepository interface {
	MigrateShapesEx() error
	CreateShapesEx([]model.ShapeEx) error
	UpdateShapesEx([]model.ShapeEx) error
	FindShapesExByTripsAndShapes() ([]model.ShapeEx, error)
	FindShapesExByTripId(tripId string) ([]model.ShapeEx, error)
	FindTripWithStopLocationByTripId(tripId string) ([]model.TripWithStopLocation, error)
}

/*
type ShapeExGeomRepository interface {
	MigrateShapesExGeom() error
	CreateShapesExGeom([]model.ShapeExGeom) error
	UpdateShapesExGeom([]model.ShapeExGeom) error
	FindShapesExGeomByTripsAndShapes() ([]model.ShapeExGeom, error)
	FindShapesExGeomByTripId(tripId string) ([]model.ShapeExGeom, error)
	FindTripWithStopLocationByTripId(tripId string) ([]model.TripWithStopLocation, error)
}
*/

type ShapeDetailRepository interface {
	MigrateShapesDetail() error
	CreateShapesDetail([]model.ShapeDetail) error
}

/*
type ShapeDetailGeomRepository interface {
	MigrateShapesDetailGeom() error
	CreateShapesDetailGeom([]model.ShapeDetailGeom) error
}
*/

type StopTimeRepository interface {
	FindStopTimesByTripId(tripId string) ([]model.StopTime, error)
}
