package ormstatic

type Route struct {
	RouteId           *string `gorm:"primaryKey;index;not null"`
	AgencyId          *string
	RouteShortName    *string
	RouteLongName     *string
	RouteDesc         *string
	RouteType         *int `gorm:"index;not null"`
	RouteUrl          *string
	RouteColor        *string
	RouteTextColor    *string
	RouteSortOrder    *int     `gorm:"index"`
	ContinuousPickup  *int     `gorm:"default:1"`
	ContinuousDropOff *int     `gorm:"default:1"`
	Trip              Trip     `gorm:"foreignKey:RouteId"`
	FareRule          FareRule `gorm:"foreignKey:RouteId"`
}

func (Route) TableName() string {
	return "routes"
}
