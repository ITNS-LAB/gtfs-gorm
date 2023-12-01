package ormstatic

type Shape struct {
	ShapeId           *string  `gorm:"primaryKey;index;not null"`
	ShapePtLat        *float64 `gorm:"not null"`
	ShapePtLon        *float64 `gorm:"not null"`
	ShapePtSequence   *int     `gorm:"primaryKey;index;not null"`
	ShapeDistTraveled *float64
}

func (Shape) TableName() string {
	return "shapes"
}
