package ormstatic

type Trip struct {
	RouteId              *string `gorm:"index;not null"`
	ServiceId            *string `gorm:"index;not null"`
	TripId               *string `gorm:"primaryKey"`
	TripHeadsign         *string
	TripShortName        *string
	DirectionId          *int        `gorm:"index"`
	BlockId              *string     `gorm:"index"`
	ShapeId              *string     `gorm:"index"`
	WheelchairAccessible *int        `gorm:"default:0"`
	BikesAllowed         *int        `gorm:"default:0"`
	StopTimes            []StopTime  `gorm:"foreignKey:TripId;references:TripId"`
	Frequencies          []Frequency `gorm:"foreignKey:TripId;references:TripId"`
}

func (Trip) TableName() string {
	return "trips"
}
