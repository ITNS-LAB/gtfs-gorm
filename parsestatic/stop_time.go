package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseStopTimes(path string) []ormstatic.StopTime {
	var stopTimes []ormstatic.StopTime
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		stopTimes = append(stopTimes, ormstatic.StopTime{
			TripId:            dataframe.IsBlank(df.GetElement("trip_id")),
			ArrivalTime:       dataframe.ParseTime(df.GetElement("arrival_time")),
			DepartureTime:     dataframe.ParseTime(df.GetElement("departure_time")),
			StopId:            dataframe.IsBlank(df.GetElement("stop_id")),
			StopSequence:      dataframe.ParseInt(df.GetElement("stop_sequence")),
			StopHeadsign:      dataframe.IsBlank(df.GetElement("stop_headsign")),
			PickupType:        dataframe.ParseEnum(df.GetElement("pickup_type")),
			DropOffType:       dataframe.ParseEnum(df.GetElement("drop_off_type")),
			ContinuousDropOff: dataframe.ParseEnum(df.GetElement("continuous_drop_off")),
			ShapeDistTraveled: dataframe.ParseFloat64(df.GetElement("shape_dist_traveled")),
			Timepoint:         dataframe.ParseEnum(df.GetElement("timepoint")),
		})
	}
	return stopTimes
}
