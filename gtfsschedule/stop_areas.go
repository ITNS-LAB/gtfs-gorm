package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type StopArea struct {
	AreaID string `gorm:"primary_key"`
	StopID string `gorm:"not null"`
}

func ParseStopArea(path string) ([]StopArea, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stop_area CSV: %w", err)
	}

	// Parse the data and create a slice of StopArea structs
	var stopAreas []StopArea
	for i := 0; i < len(df.Records); i++ {
		areaID, err := df.GetString(i, "area_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'area_id' at row %d: %w", i, err)
		}

		stopID, err := df.GetString(i, "stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_id' at row %d: %w", i, err)
		}

		// Create the StopArea struct and append to the list
		stopAreas = append(stopAreas, StopArea{
			AreaID: areaID,
			StopID: stopID,
		})
	}

	return stopAreas, nil
}
