package ormstatic

import geomdatatypes "github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"

type Shape struct {
	ShapeId           *string  `gorm:"primaryKey"`
	ShapePtLat        *float64 `gorm:"not null"`
	ShapePtLon        *float64 `gorm:"not null"`
	ShapePtSequence   *int     `gorm:"primaryKey"`
	ShapeDistTraveled *float64
	Geom              *geomdatatypes.Geometry `gorm:"index"`
}

func (Shape) TableName() string {
	return "shapes"
}
