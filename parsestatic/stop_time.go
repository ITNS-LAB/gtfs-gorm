package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseStopTimes(path string) ([]gtfsjp.StopTime, error) {
	var stopTimes []gtfsjp.StopTime
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

		stopTimes = append(stopTimes, gtfsjp.StopTime{
			TripId:            util.IsBlank(df.GetElement("trip_id")),
			ArrivalTime:       util.ParseTime(df.GetElement("arrival_time")),
			DepartureTime:     util.ParseTime(df.GetElement("departure_time")),
			StopId:            util.IsBlank(df.GetElement("stop_id")),
			StopSequence:      util.ParseInt(df.GetElement("stop_sequence")),
			StopHeadsign:      util.IsBlank(df.GetElement("stop_headsign")),
			PickupType:        util.ParseEnum(df.GetElement("pickup_type")),
			DropOffType:       util.ParseEnum(df.GetElement("drop_off_type")),
			ContinuousDropOff: util.ParseEnum(df.GetElement("continuous_drop_off")),
			ShapeDistTraveled: util.ParseFloat64(df.GetElement("shape_dist_traveled")),
			Timepoint:         util.ParseEnum(df.GetElement("timepoint")),
		})
	}
	return stopTimes, nil
}
