package gtfsschedule

type Attribution struct {
	AttributionID    *string
	AgencyID         *string
	RouteID          *string
	TripID           *string
	OrganizationName string `gorm:"not null"`
	IsProducer       *bool
	IsOperator       *bool
	IsAuthority      *bool
	AttributionURL   *string
	AttributionEmail *string
	AttributionPhone *string
}
