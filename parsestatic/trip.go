package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
)

func ParseTrips(path string) ([]ormstatic.Trip, error) {
	var trips []ormstatic.Trip
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return trips, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		trips = append(trips, ormstatic.Trip{
			RouteId:              util.IsBlank(df.GetElement("route_id")),
			ServiceId:            util.IsBlank(df.GetElement("service_id")),
			TripId:               util.IsBlank(df.GetElement("trip_id")),
			TripHeadsign:         util.IsBlank(df.GetElement("trip_headsign")),
			TripShortName:        util.IsBlank(df.GetElement("trip_short_name")),
			DirectionId:          util.ParseEnum(df.GetElement("direction_id")),
			BlockId:              util.IsBlank(df.GetElement("block_id")),
			ShapeId:              util.IsBlank(df.GetElement("shape_id")),
			WheelchairAccessible: util.ParseEnum(df.GetElement("wheelchair_accessible")),
			BikesAllowed:         util.ParseEnum(df.GetElement("bikes_allowed")),
		})
	}
	return trips, nil
}
