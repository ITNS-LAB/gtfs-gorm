package static

type Stop struct {
	StopId             *string `gorm:"primaryKey;index;not null"`
	StopCode           *string
	StopName           *string
	StopDesc           *string
	StopLat            *float64
	StopLon            *float64
	ZoneId             *string `gorm:"unique"`
	StopUrl            *string
	LocationType       *int `gorm:"default:0"`
	ParentStation      *string
	StopTimezone       *string
	WheelchairBoarding *int `gorm:"default:0"`
	LevelId            *string
	PlatformCode       *string
	StopTime           StopTime `gorm:"foreignKey:StopId"`
}

func (Stop) TableName() string {
	return "stops"
}
