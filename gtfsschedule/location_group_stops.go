package gtfsschedule

type LocationGroupStop struct {
	LocationGroupID string `gorm:"not null"`
	StopID          string `gorm:"not null"`
}
