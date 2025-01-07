package gtfsjp

import (
	"database/sql"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
)

type ShapeEx struct {
	TripId            string  `gorm:"primaryKey"`
	ShapeId           string  `gorm:"primaryKey"`
	ShapePtLat        float64 `gorm:"not null"`
	ShapePtLon        float64 `gorm:"not null"`
	ShapePtSequence   int     `gorm:"primaryKey"`
	ShapeDistTraveled float64
	StopId            sql.NullString
	Stop              Stop `gorm:"foreignKey:StopId"`
}

func (ShapeEx) TableName() string {
	return "shapes_ex"
}

type ShapeExGeom struct {
	TripId            string  `gorm:"primaryKey"`
	ShapeId           string  `gorm:"primaryKey"`
	ShapePtLat        float64 `gorm:"not null"`
	ShapePtLon        float64 `gorm:"not null"`
	ShapePtSequence   int     `gorm:"primaryKey"`
	ShapeDistTraveled float64
	StopId            sql.NullString
	Geom              gormdatatypes.Geometry `gorm:"index"`
	Stop              Stop                   `gorm:"foreignKey:StopId"`
}

func (ShapeExGeom) TableName() string {
	return "shapes_ex"
}
