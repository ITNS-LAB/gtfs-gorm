package parsestatic

import (
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
			return []ormstatic.Route{}, err
		}

		routeId, err := dataframe.ParseString(df.GetElement("route_id"))
		if err != nil {
			return []ormstatic.Route{}, err
		}

		routeType, err := dataframe.ParseInt16(df.GetElement("route_type"))
		if err != nil {
			return []ormstatic.Route{}, err
		}

		routeSortOrder, err := dataframe.ParseNullInt64(df.GetElement("route_sort_order"))
		continuousPickup, err := dataframe.ParseNullInt16(df.GetElement("continuous_pickup"))
		continuousDropOff, err := dataframe.ParseNullInt16(df.GetElement("continuous_drop_off"))

		routes = append(routes, ormstatic.Route{
			RouteId:           routeId,
			AgencyId:          df.GetElement("agency_id"),
			RouteShortName:    df.GetElement("route_short_name"),
			RouteLongName:     df.GetElement("route_long_name"),
			RouteDesc:         df.GetElement("route_desc"),
			RouteType:         routeType,
			RouteUrl:          df.GetElement("route_url"),
			RouteColor:        df.GetElement("route_color"),
			RouteTextColor:    df.GetElement("route_text_color"),
			RouteSortOrder:    routeSortOrder,
			ContinuousPickup:  continuousPickup,
			ContinuousDropOff: continuousDropOff,
		})
	}
	return routes, nil
}
