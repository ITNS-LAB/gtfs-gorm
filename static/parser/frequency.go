package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseFrequencies(path string) []orm.Frequency {
	var frequencies []orm.Frequency
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		frequencies = append(frequencies, orm.Frequency{
			TripId:      dataframe.IsBlank(df.GetElement("trip_id")),
			StartTime:   dataframe.ParseTime(df.GetElement("start_time")),
			EndTime:     dataframe.ParseTime(df.GetElement("end_time")),
			HeadwaySecs: dataframe.ParseInt(df.GetElement("headway_secs")),
			ExactTimes:  dataframe.ParseEnum(df.GetElement("exact_times")),
		})
	}
	return frequencies
}
