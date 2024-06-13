package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
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
			return []ormstatic.Frequency{}, err
		}

		tripId, err := dataframe.ParseString(df.GetElement("trip_id"))
		if err != nil {
			return []ormstatic.Frequency{}, err
		}

		startTime, err := dataframe.ParseDataTypesTime(df.GetElement("start_time"))
		if err != nil {
			return []ormstatic.Frequency{}, err
		}

		endTime, err := dataframe.ParseDataTypesTime(df.GetElement("end_time"))
		if err != nil {
			return []ormstatic.Frequency{}, err
		}

		headwaySecs, err := dataframe.ParseInt32(df.GetElement("headway_secs"))
		if err != nil {
			return []ormstatic.Frequency{}, err
		}

		exactTimes, err := dataframe.ParseNullInt16(df.GetElement("exact_times"))

		frequencies = append(frequencies, ormstatic.Frequency{
			TripId:      tripId,
			StartTime:   startTime,
			EndTime:     endTime,
			HeadwaySecs: headwaySecs,
			ExactTimes:  exactTimes,
		})
	}
	return frequencies, nil
}
