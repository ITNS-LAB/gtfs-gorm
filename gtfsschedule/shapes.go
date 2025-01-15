package gtfsschedule

type Shape struct {
	ShapeID           string  `gorm:"primary_key"`
	ShapePtLat        float64 `gorm:"not null"`
	ShapePtLon        float64 `gorm:"not null"`
	ShapePtSequence   int     `gorm:"not null"`
	ShapeDistTraveled *float64
}
