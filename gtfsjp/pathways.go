package gtfsjp

type Pathway struct {
	PathwayId            *string `gorm:"primaryKey"`
	FromStopId           *string `gorm:"index;not null"`
	ToStopId             *string `gorm:"index;not null"`
	PathwayMode          *int    `gorm:"index;not null"`
	IsBidirectional      *int    `gorm:"index;not null"`
	Length               *float64
	TraversalTime        *int
	StairCount           *int
	MaxSlope             *float64 `gorm:"default:0"`
	MinWidth             *float64
	SignpostedAs         *string
	ReversedSignpostedAs *string
}

func (Pathway) TableName() string {
	return "pathways"
}
