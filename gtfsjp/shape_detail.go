package gtfsjp

import (
	"database/sql"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"gorm.io/datatypes"
	"time"
)

type ShapeDetail struct {
	ShapeId               string  `gorm:"primaryKey"`
	ShapePtLat            float64 `gorm:"not null"`
	ShapePtLon            float64 `gorm:"not null"`
	ShapeDetailPtSequence int     `gorm:"primaryKey"`
	ShapeDistTraveled     float64
}

func (ShapeDetail) TableName() string {
	return "shapes_detail"
}

type ShapeDetailEx struct {
	TripId                string  `gorm:"primaryKey"`
	ShapeId               string  `gorm:"primaryKey"`
	ShapePtLat            float64 `gorm:"not null"`
	ShapePtLon            float64 `gorm:"not null"`
	ShapeDetailPtSequence int     `gorm:"primaryKey"`
	ShapeDistTraveled     float64
	StopId                sql.NullString
	ShapesTime            datatypes.Time
}

func (ShapeDetailEx) TableName() string { return "shapes_detail_ex" }

type ShapeDetailExTemp struct {
	TripId                string  `gorm:"primaryKey"`
	ShapeId               string  `gorm:"primaryKey"`
	ShapePtLat            float64 `gorm:"not null"`
	ShapePtLon            float64 `gorm:"not null"`
	ShapeDetailPtSequence int     `gorm:"primaryKey"`
	ShapeDistTraveled     float64
	StopId                sql.NullString
	ShapesTime            time.Time
}

func (ShapeDetailExTemp) TableName() string { return "shapes_detail_ex_temp" }

type ShapeDetailGeom struct {
	ShapeId               string  `gorm:"primaryKey"`
	ShapePtLat            float64 `gorm:"not null"`
	ShapePtLon            float64 `gorm:"not null"`
	ShapeDetailPtSequence int     `gorm:"primaryKey"`
	ShapeDistTraveled     float64
	Geom                  gormdatatypes.Geometry `gorm:"index"`
}

func (ShapeDetailGeom) TableName() string {
	return "shapes_detail"
}
