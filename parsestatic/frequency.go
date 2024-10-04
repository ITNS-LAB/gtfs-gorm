package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
)

func ParseFrequencies(path string) ([]ormstatic.Frequency, error) {
	var frequencies []ormstatic.Frequency
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return frequencies, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		frequencies = append(frequencies, ormstatic.Frequency{
			TripId:      util.IsBlank(df.GetElement("trip_id")),
			StartTime:   util.ParseTime(df.GetElement("start_time")),
			EndTime:     util.ParseTime(df.GetElement("end_time")),
			HeadwaySecs: util.ParseInt(df.GetElement("headway_secs")),
			ExactTimes:  util.ParseEnum(df.GetElement("exact_times")),
		})
	}
	return frequencies, nil
}
