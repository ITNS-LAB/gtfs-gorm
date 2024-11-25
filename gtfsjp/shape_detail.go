package gtfsjp

import geomdatatypes "github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"

type ShapeDetail struct {
	ShapeId               *string  `gorm:"primaryKey"`
	ShapePtLat            *float64 `gorm:"not null"`
	ShapePtLon            *float64 `gorm:"not null"`
	ShapeDetailPtSequence *int     `gorm:"primaryKey"`
	ShapeDistTraveled     *float64
	Geom                  *geomdatatypes.Geometry `gorm:"index"`
}

func (ShapeDetail) TableName() string {
	return "shapes_detail"
}
