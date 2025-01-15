package gtfsschedule

type Pathway struct {
	PathwayID            string `gorm:"primaryKey"`
	FromStopID           string `gorm:"not null"`
	ToStopID             string `gorm:"not null"`
	PathwayMode          int    `gorm:"not null"`
	IsBidirectional      int    `gorm:"not null"`
	Length               *float64
	TraversalTime        *int
	StairCount           *int
	MaxSlope             *float64
	MinWidth             *float64
	SignpostedAs         *string
	ReversedSignpostedAs *string
}
