package ormstatic

import "database/sql"

type Shape struct {
	ShapeId           string  `gorm:"primaryKey;index;not null"`
	ShapePtLat        float64 `gorm:"not null"`
	ShapePtLon        float64 `gorm:"not null"`
	ShapePtSequence   int     `gorm:"primaryKey;index;not null"`
	ShapeDistTraveled sql.NullFloat64
}

func (Shape) TableName() string {
	return "shapes"
}

func (s Shape) GetShapeId() any {
	return s.ShapeId
}

func (s Shape) GetShapePtLat() any {
	return s.ShapePtLat
}

func (s Shape) GetShapePtLon() any {
	return s.ShapePtLon
}

func (s Shape) GetShapePtSequence() any {
	return s.ShapePtSequence
}

func (s Shape) GetShapeDistTraveled() any {
	if s.ShapeDistTraveled.Valid {
		return s.ShapeDistTraveled.Float64
	}
	return nil
}
