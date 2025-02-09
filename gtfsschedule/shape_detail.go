package gtfsschedule

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

/*
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

*/
