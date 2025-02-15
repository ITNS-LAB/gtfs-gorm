package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type LocationGroupStop struct {
	LocationGroupId string `gorm:"not null"`
	StopId          string `gorm:"not null"`
}

func ParseLocationGroupStop(path string) ([]LocationGroupStop, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open location_group_stop CSV: %w", err)
	}

	// Parse the CSV data into a slice of LocationGroupStop structs
	var locationGroupStops []LocationGroupStop
	for i := 0; i < len(df.Records); i++ {
		locationGroupID, err := df.GetString(i, "location_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'location_group_id' at row %d: %w", i, err)
		}

		stopID, err := df.GetString(i, "stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_id' at row %d: %w", i, err)
		}

		// Create LocationGroupStop struct and append to the list
		locationGroupStops = append(locationGroupStops, LocationGroupStop{
			LocationGroupId: locationGroupID,
			StopId:          stopID,
		})
	}

	return locationGroupStops, nil
}
