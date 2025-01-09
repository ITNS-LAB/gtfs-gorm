package gtfsschedule

type Route struct {
	RouteId           string `gorm:"primary_key"`
	AgencyId          *string
	RouteShortName    *string
	RouteLongName     *string
	RouteDesc         *string
	RouteType         int `gorm:"not null"`
	RouteUrl          *string
	RouteColor        *string
	RouteTextColor    *string
	RouteSortOrder    *string
	ContinuousPickup  *int
	ContinuousDropOff *int
	NetworkId         *string
}
