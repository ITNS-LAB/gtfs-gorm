package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Route struct {
	RouteId           string `gorm:"primaryKey"`
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
	NetworkId         string      `gorm:"index;unique"`
	Trips             []Trips     `gorm:"foreignKey:RouteId;references:RouteId"`
	FareRules         []FareRules `gorm:"foreignKey:RouteId;references:RouteId"`
	//FareLeg                       FareLeg          `gorm:"foreignKey:NetworkId;references:NetworkId"`
	FareLegJoinRulesFromNetworkID FareLegJoinRules `gorm:"foreignKey:FromNetworkId;references:NetworkId"`
	FareLegJoinRulesToStopID      FareLegJoinRules `gorm:"foreignKey:ToStopId;references:NetworkId"`
	RouteNetwork                  []RouteNetwork   `gorm:"foreignKey:RouteId;references:RouteId"`
	TransferFromRouteID           []Transfer       `gorm:"foreignKey:FromRouteId;references:RouteId"`
	TransferToRouteID             []Transfer       `gorm:"foreignKey:ToRouteId;references:RouteId"`
	Attribution                   []Attribution    `gorm:"foreignKey:RouteId;references:RouteId"`
}

func ParseRoutes(path string) ([]Route, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open Routes CSV: %w", err)
	}

	// データを解析してRoute構造体のスライスを作成
	var routes []Route
	for i := 0; i < len(df.Records); i++ {
		routeId, err := df.GetString(i, "route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_id' at row %d: %w", i, err)
		}

		agencyId, err := df.GetStringPtr(i, "agency_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_id' at row %d: %w", i, err)
		}

		routeShortName, err := df.GetStringPtr(i, "route_short_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_short_name' at row %d: %w", i, err)
		}

		routeLongName, err := df.GetStringPtr(i, "route_long_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_long_name' at row %d: %w", i, err)
		}

		routeDesc, err := df.GetStringPtr(i, "route_desc")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_desc' at row %d: %w", i, err)
		}

		routeType, err := df.GetInt(i, "route_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_type' at row %d: %w", i, err)
		}

		routeUrl, err := df.GetStringPtr(i, "route_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_url' at row %d: %w", i, err)
		}

		routeColor, err := df.GetStringPtr(i, "route_color")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_color' at row %d: %w", i, err)
		}

		routeTextColor, err := df.GetStringPtr(i, "route_text_color")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_text_color' at row %d: %w", i, err)
		}

		routeSortOrder, err := df.GetStringPtr(i, "route_sort_order")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_sort_order' at row %d: %w", i, err)
		}

		continuousPickup, err := df.GetIntPtr(i, "continuous_pickup")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'continuous_pickup' at row %d: %w", i, err)
		}

		continuousDropOff, err := df.GetIntPtr(i, "continuous_drop_off")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'continuous_drop_off' at row %d: %w", i, err)
		}

		networkId, err := df.GetString(i, "network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'network_id' at row %d: %w", i, err)
		}

		// Route構造体を作成しリストに追加
		routes = append(routes, Route{
			RouteId:           routeId,
			AgencyId:          agencyId,
			RouteShortName:    routeShortName,
			RouteLongName:     routeLongName,
			RouteDesc:         routeDesc,
			RouteType:         routeType,
			RouteUrl:          routeUrl,
			RouteColor:        routeColor,
			RouteTextColor:    routeTextColor,
			RouteSortOrder:    routeSortOrder,
			ContinuousPickup:  continuousPickup,
			ContinuousDropOff: continuousDropOff,
			NetworkId:         networkId,
		})
	}

	return routes, nil
}
