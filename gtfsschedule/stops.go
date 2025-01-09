package gtfsschedule

type Stop struct {
	stopId             string `gorm:"primary_key"`
	stopCode           *string
	stopName           *string
	ttsStopName        *string
	stopDesc           *string
	stopLat            float64 `gorm:"not null"`
	stopLon            float64 `gorm:"not null"`
	zoneId             *string
	stopUrl            *string
	locationType       *string
	parentStation      int `gorm:"not null"`
	stopTimezone       *string
	wheelchairBoarding *string
	levelId            *string
	platformCode       *string
}
