package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseTrips(path string) []static.Trip {
	var trips []static.Trip
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		trips = append(trips, static.Trip{
			RouteId:              dataframe.IsBlank(df.GetElement("route_id")),
			ServiceId:            dataframe.IsBlank(df.GetElement("service_id")),
			TripId:               dataframe.IsBlank(df.GetElement("trip_id")),
			TripHeadsign:         dataframe.IsBlank(df.GetElement("trip_headsign")),
			TripShortName:        dataframe.IsBlank(df.GetElement("trip_short_name")),
			DirectionId:          dataframe.ParseEnum(df.GetElement("direction_id")),
			BlockId:              dataframe.IsBlank(df.GetElement("block_id")),
			ShapeId:              dataframe.IsBlank(df.GetElement("shape_id")),
			WheelchairAccessible: dataframe.ParseEnum(df.GetElement("wheelchair_accessible")),
			BikesAllowed:         dataframe.ParseEnum(df.GetElement("bikes_allowed")),
		})
	}
	return trips
}
