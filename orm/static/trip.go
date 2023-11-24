package static

type Trip struct {
	RouteId              *string `gorm:"index;not null"`
	ServiceId            *string `gorm:"index;not null"`
	TripId               *string `gorm:"primaryKey;index;not null"`
	TripHeadsign         *string
	TripShortName        *string
	DirectionId          *int      `gorm:"index"`
	BlockId              *string   `gorm:"index"`
	ShapeId              *string   `gorm:"index"`
	WheelchairAccessible *int      `gorm:"default:0"`
	BikesAllowed         *int      `gorm:"default:0"`
	Frequency            Frequency `gorm:"foreignKey:TripId"`
}

func (Trip) TableName() string {
	return "trips"
}
