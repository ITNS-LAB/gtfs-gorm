package parsestatic

import (
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
			return []ormstatic.Trip{}, err
		}

		routeId, err := dataframe.ParseString(df.GetElement("route_id"))
		if err != nil {
			return []ormstatic.Trip{}, err
		}

		serviceId, err := dataframe.ParseString(df.GetElement("service_id"))
		if err != nil {
			return []ormstatic.Trip{}, err
		}

		tripId, err := dataframe.ParseString(df.GetElement("trip_id"))
		if err != nil {
			return []ormstatic.Trip{}, err
		}

		directionId, err := dataframe.ParseNullInt16(df.GetElement("direction_id"))
		wheelchairAccessible, err := dataframe.ParseNullInt16(df.GetElement("wheelchair_accessible"))
		bikesAllowed, err := dataframe.ParseNullInt16(df.GetElement("bikes_allowed"))

		trips = append(trips, ormstatic.Trip{
			RouteId:              routeId,
			ServiceId:            serviceId,
			TripId:               tripId,
			TripHeadsign:         df.GetElement("trip_headsign"),
			TripShortName:        df.GetElement("trip_short_name"),
			DirectionId:          directionId,
			BlockId:              df.GetElement("block_id"),
			ShapeId:              df.GetElement("shape_id"),
			WheelchairAccessible: wheelchairAccessible,
			BikesAllowed:         bikesAllowed,
		})
	}
	return trips, nil
}
