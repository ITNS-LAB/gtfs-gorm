package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Route struct {
	RouteId         string `gorm:"primaryKey"`
	AgencyId        string // primaryにするか検討中
	RouteShortName  *string
	RouteLongName   *string
	RouteDesc       *string
	RouteType       int `gorm:"not null"`
	RouteUrl        *string
	RouteColor      *string
	RouteTextColor  *string
	JpParentRouteId *string
	Trips           []Trip     `gorm:"foreignKey:RouteId;references:RouteId"`
	FareRules       []FareRule `gorm:"foreignKey:RouteId;references:RouteId"`
}

func (Route) TableName() string {
	return "routes"
}

func ParseRoutes(path string) ([]Route, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open routes CSV: %w", err)
	}

	// データを解析して Route 構造体のスライスを作成
	var routes []Route
	for i := 0; i < len(df.Records); i++ {
		routeId, err := df.GetString(i, "route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_id' at row %d: %w", i, err)
		}

		agencyId, err := df.GetString(i, "agency_id")
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

		jpParentRouteId, err := df.GetStringPtr(i, "jp_parent_route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_parent_route_id' at row %d: %w", i, err)
		}

		// Route 構造体を作成しリストに追加
		routes = append(routes, Route{
			RouteId:         routeId,
			AgencyId:        agencyId,
			RouteShortName:  routeShortName,
			RouteLongName:   routeLongName,
			RouteDesc:       routeDesc,
			RouteType:       routeType,
			RouteUrl:        routeUrl,
			RouteColor:      routeColor,
			RouteTextColor:  routeTextColor,
			JpParentRouteId: jpParentRouteId,
		})
	}

	return routes, nil
}
