package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
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
	return trips, nil
}
