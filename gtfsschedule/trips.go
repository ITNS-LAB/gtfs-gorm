package gtfsschedule

type Trips struct {
	RouteId              string `gorm:"index;not null"`
	ServiceId            string `gorm:"index;not null"`
	TripId               string `gorm:"primaryKey"`
	TripHeadsign         *string
	TripShortName        *string
	DirectionId          *int    `gorm:"index"`
	BlockId              *string `gorm:"index"`
	ShapeId              *string `gorm:"index"`
	WheelchairAccessible *int
	BikesAllowed         *int
}
