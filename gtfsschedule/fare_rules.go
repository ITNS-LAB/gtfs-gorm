package gtfsschedule

type FareRules struct {
	FareID        int `gorm:"primary_key"`
	RouteID       *string
	OriginID      *string
	DestinationID *string
	ContainsId    *string
}
