package gtfsschedule

type Route struct {
	RouteId                       string `gorm:"primary_key"`
	AgencyId                      *string
	RouteShortName                *string
	RouteLongName                 *string
	RouteDesc                     *string
	RouteType                     int `gorm:"not null"`
	RouteUrl                      *string
	RouteColor                    *string
	RouteTextColor                *string
	RouteSortOrder                *string
	ContinuousPickup              *int
	ContinuousDropOff             *int
	NetworkId                     *string
	Trips                         []Trips          `gorm:"foreignKey:RouteId;references:RouteId "`
	FareRules                     []FareRules      `gorm:"foreignKey:RouteId;references:RouteId "`
	FareLeg                       FareLeg          `gorm:"foreignKey:NetworkId;references:NetworkId "`
	FareLegJoinRulesFromNetworkID FareLegJoinRules `gorm:"foreignKey:NetworkId;references:FromNetworkID "`
	FareLegJoinRulesToStopID      FareLegJoinRules `gorm:"foreignKey:NetworkId;references:ToStopID "`
	RouteNetwork                  []RouteNetwork   `gorm:"foreignKey:RouteId;references:RouteID "`
	TransferFromRouteID           []Transfer       `gorm:"foreignKey:RouteId;references:FromRouteID "`
	TransferToRouteID             []Transfer       `gorm:"foreignKey:RouteId;references:ToRouteID "`
	Attribution                   []Attribution    `gorm:"foreignKey:RouteId;references:RouteId "`
}
