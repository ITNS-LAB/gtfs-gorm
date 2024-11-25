package gtfsjp

type Route struct {
	RouteId           *string `gorm:"primaryKey"`
	AgencyId          *string
	RouteShortName    *string
	RouteLongName     *string
	RouteDesc         *string
	RouteType         *int `gorm:"index;not null"`
	RouteUrl          *string
	RouteColor        *string
	RouteTextColor    *string
	RouteSortOrder    *int       `gorm:"index"`
	ContinuousPickup  *int       `gorm:"default:1"`
	ContinuousDropOff *int       `gorm:"default:1"`
	Trips             []Trip     `gorm:"foreignKey:RouteId;references:RouteId"`
	FareRules         []FareRule `gorm:"foreignKey:RouteId;references:RouteId"`
}

func (Route) TableName() string {
	return "routes"
}
