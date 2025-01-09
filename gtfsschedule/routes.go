package gtfsschedule

type Route struct {
	RouteId           string `gorm:"primary_key"`
	AgencyId          *string
	RouteShortName    *string
	RouteLongName     *string
	RouteDesc         *string
	RouteType         string `gorm:"not null"`
	RouteUrl          *string
	RouteColor        *string
	RouteTextColor    *string
	RouteSortOrder    *string
	ContinuousPickup  *string
	ContinuousDropOff *string
	NetworkId         *string
}
