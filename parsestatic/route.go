package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseRoutes(path string) ([]gtfsjp.Route, error) {
	var routes []gtfsjp.Route
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

		routes = append(routes, gtfsjp.Route{
			RouteId:           util.IsBlank(df.GetElement("route_id")),
			AgencyId:          util.IsBlank(df.GetElement("agency_id")),
			RouteShortName:    util.IsBlank(df.GetElement("route_short_name")),
			RouteLongName:     util.IsBlank(df.GetElement("route_long_name")),
			RouteDesc:         util.IsBlank(df.GetElement("route_desc")),
			RouteType:         util.ParseEnum(df.GetElement("route_type")),
			RouteUrl:          util.IsBlank(df.GetElement("route_url")),
			RouteColor:        util.IsBlank(df.GetElement("route_color")),
			RouteTextColor:    util.IsBlank(df.GetElement("route_text_color")),
			RouteSortOrder:    util.ParseInt(df.GetElement("route_sort_order")),
			ContinuousPickup:  util.ParseEnum(df.GetElement("continuous_pickup")),
			ContinuousDropOff: util.ParseEnum(df.GetElement("continuous_drop_off")),
		})
	}
	return routes, nil
}
