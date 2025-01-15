package gtfsschedule

type Frequencies struct {
	TripID      string `gorm:"primary_key"`
	StartTime   string `gorm:"not null"`
	EndTime     string `gorm:"not null"`
	HeadwaySecs int    `gorm:"not null"`
	ExactTimes  *int
}
