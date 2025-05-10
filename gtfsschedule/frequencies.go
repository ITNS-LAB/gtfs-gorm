package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Frequencies struct {
	TripId      string `gorm:"primaryKey"`
	StartTime   string `gorm:"not null"`
	EndTime     string `gorm:"not null"`
	HeadwaySecs int    `gorm:"not null"`
	ExactTimes  *int
}

func (Frequencies) TableName() string {
	return "frequencies"
}

func ParseFrequencies(path string) ([]Frequencies, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open frequencies CSV: %w", err)
	}

	// Parse the CSV data into a slice of Frequencies structs
	var frequencies []Frequencies
	for i := 0; i < len(df.Records); i++ {
		tripID, err := df.GetString(i, "trip_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_id' at row %d: %w", i, err)
		}

		startTime, err := df.GetString(i, "start_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'start_time' at row %d: %w", i, err)
		}

		endTime, err := df.GetString(i, "end_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'end_time' at row %d: %w", i, err)
		}

		headwaySecs, err := df.GetInt(i, "headway_secs")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'headway_secs' at row %d: %w", i, err)
		}

		exactTimes, err := df.GetIntPtr(i, "exact_times")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'exact_times' at row %d: %w", i, err)
		}

		// Create Frequencies struct and append to the list
		frequencies = append(frequencies, Frequencies{
			TripId:      tripID,
			StartTime:   startTime,
			EndTime:     endTime,
			HeadwaySecs: headwaySecs,
			ExactTimes:  exactTimes,
		})
	}

	return frequencies, nil
}

type FrequenciesGeom struct {
	TripId      string `gorm:"primaryKey"`
	StartTime   string `gorm:"not null"`
	EndTime     string `gorm:"not null"`
	HeadwaySecs int    `gorm:"not null"`
	ExactTimes  *int
}

func (FrequenciesGeom) TableName() string {
	return "frequencies"
}

func ParseFrequenciesGeom(path string) ([]FrequenciesGeom, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open frequencies CSV: %w", err)
	}

	// Parse the CSV data into a slice of Frequencies structs
	var frequencies []FrequenciesGeom
	for i := 0; i < len(df.Records); i++ {
		tripID, err := df.GetString(i, "trip_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_id' at row %d: %w", i, err)
		}

		startTime, err := df.GetString(i, "start_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'start_time' at row %d: %w", i, err)
		}

		endTime, err := df.GetString(i, "end_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'end_time' at row %d: %w", i, err)
		}

		headwaySecs, err := df.GetInt(i, "headway_secs")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'headway_secs' at row %d: %w", i, err)
		}

		exactTimes, err := df.GetIntPtr(i, "exact_times")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'exact_times' at row %d: %w", i, err)
		}

		// Create Frequencies struct and append to the list
		frequencies = append(frequencies, FrequenciesGeom{
			TripId:      tripID,
			StartTime:   startTime,
			EndTime:     endTime,
			HeadwaySecs: headwaySecs,
			ExactTimes:  exactTimes,
		})
	}

	return frequencies, nil
}
