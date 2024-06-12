package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseStopTimes(path string) ([]ormstatic.StopTime, error) {
	var stopTimes []ormstatic.StopTime
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return stopTimes, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		tripId, err := dataframe.ParseString(df.GetElement("trip_id"))
		if err != nil {
			return []ormstatic.StopTime{}, err
		}
		arrivalTime, err := dataframe.ParseNullDataTypesTime(df.GetElement("arrival_time"))
		departureTime, err := dataframe.ParseNullDataTypesTime(df.GetElement("departure_time"))
		stopId, err := dataframe.ParseString(df.GetElement("stop_id"))
		if err != nil {
			return []ormstatic.StopTime{}, err
		}
		stopSequence, err := dataframe.ParseInt64(df.GetElement("stop_sequence"))
		if err != nil {
			return []ormstatic.StopTime{}, err
		}
		pickupType, err := dataframe.ParseNullInt16(df.GetElement("pickup_type"))
		dropOffType, err := dataframe.ParseNullInt16(df.GetElement("drop_off_type"))
		continuousPickup, err := dataframe.ParseNullInt16(df.GetElement("continuous_pickup"))
		continuousDropOff, err := dataframe.ParseNullInt16(df.GetElement("continuous_drop_off"))
		shapeDistTraveled, err := dataframe.ParseNullFloat64(df.GetElement("shape_dist_traveled"))
		timePoint, err := dataframe.ParseNullInt16(df.GetElement("timepoint"))

		stopTimes = append(stopTimes, ormstatic.StopTime{
			TripId:            tripId,
			ArrivalTime:       arrivalTime,
			DepartureTime:     departureTime,
			StopId:            stopId,
			StopSequence:      stopSequence,
			StopHeadsign:      df.GetElement("stop_headsign"),
			PickupType:        pickupType,
			DropOffType:       dropOffType,
			ContinuousPickup:  continuousPickup,
			ContinuousDropOff: continuousDropOff,
			ShapeDistTraveled: shapeDistTraveled,
			Timepoint:         timePoint,
		})
	}
	return stopTimes, nil
}
