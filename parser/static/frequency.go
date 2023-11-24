package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseFrequencies(path string) []static.Frequency {
	var frequencies []static.Frequency
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		frequencies = append(frequencies, static.Frequency{
			TripId:      dataframe.IsBlank(df.GetElement("trip_id")),
			StartTime:   dataframe.ParseTime(df.GetElement("start_time")),
			EndTime:     dataframe.ParseTime(df.GetElement("end_time")),
			HeadwaySecs: dataframe.ParseInt(df.GetElement("headway_secs")),
			ExactTimes:  dataframe.ParseEnum(df.GetElement("exact_times")),
		})
	}
	return frequencies
}
