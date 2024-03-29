package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseRoutes(path string) ([]ormstatic.Route, error) {
	var routes []ormstatic.Route
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return routes, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		routes = append(routes, ormstatic.Route{
			RouteId:           dataframe.IsBlank(df.GetElement("route_id")),
			AgencyId:          dataframe.IsBlank(df.GetElement("agency_id")),
			RouteShortName:    dataframe.IsBlank(df.GetElement("route_short_name")),
			RouteLongName:     dataframe.IsBlank(df.GetElement("route_long_name")),
			RouteDesc:         dataframe.IsBlank(df.GetElement("route_desc")),
			RouteType:         dataframe.ParseEnum(df.GetElement("route_type")),
			RouteUrl:          dataframe.IsBlank(df.GetElement("route_url")),
			RouteColor:        dataframe.IsBlank(df.GetElement("route_color")),
			RouteTextColor:    dataframe.IsBlank(df.GetElement("route_text_color")),
			RouteSortOrder:    dataframe.ParseInt(df.GetElement("route_sort_order")),
			ContinuousPickup:  dataframe.ParseEnum(df.GetElement("continuous_pickup")),
			ContinuousDropOff: dataframe.ParseEnum(df.GetElement("continuous_drop_off")),
		})
	}
	return routes, nil
}
