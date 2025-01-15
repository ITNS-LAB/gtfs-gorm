package gtfsschedule

type FareLeg struct {
	LegGroupID           *string
	NetworkID            *string
	FromAreaID           *string
	ToAreaID             *string
	FromTimeframeGroupID *string
	ToTimeframeGroupID   *string
	FareProductID        string `gorm:"not null"`
	RulePriority         *int
}
