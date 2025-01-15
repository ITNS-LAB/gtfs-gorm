package gtfsschedule

type FareLegJoinRules struct {
	FromNetworkID string `gorm:"primary_key"`
	ToNetworkID   string `gorm:"not null"`
	FromStopID    *string
	ToStopID      *string
}
