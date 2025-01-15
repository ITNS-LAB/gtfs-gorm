package gtfsschedule

type StopArea struct {
	AreaID string `gorm:"primary_key"`
	StopID string `gorm:"not null"`
}
