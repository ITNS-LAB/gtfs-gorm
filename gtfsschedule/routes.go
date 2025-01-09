package gtfsschedule

type Route struct {
	routeId           string `gorm:"primary_key"`
	agencyId          *string
	routeShortName    *string
	routeLongName     *string
	routeDesc         *string
	routeType         string `gorm:"not null"`
	routeUrl          *string
	routeColor        *string
	routeTextColor    *string
	routeSortOrder    *string
	continuousPickup  *string
	continuousDropOff *string
	networkId         *string
}
