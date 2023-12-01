package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseFrequencies(path string) []ormstatic.Frequency {
	var frequencies []ormstatic.Frequency
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		frequencies = append(frequencies, ormstatic.Frequency{
			TripId:      dataframe.IsBlank(df.GetElement("trip_id")),
			StartTime:   dataframe.ParseTime(df.GetElement("start_time")),
			EndTime:     dataframe.ParseTime(df.GetElement("end_time")),
			HeadwaySecs: dataframe.ParseInt(df.GetElement("headway_secs")),
			ExactTimes:  dataframe.ParseEnum(df.GetElement("exact_times")),
		})
	}
	return frequencies
}
