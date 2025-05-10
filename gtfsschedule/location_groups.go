package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type LocationGroup struct {
	LocationGroupId   string `gorm:"primaryKey"`
	LocationGroupName *string
	StopTimes         []StopTimes         `gorm:"foreignKey:LocationGroupId;references:LocationGroupId"`
	LocationGroupStop []LocationGroupStop `gorm:"foreignKey:LocationGroupId;references:LocationGroupId"`
}

func (LocationGroup) TableName() string {
	return "location_group"
}

func ParseLocationGroup(path string) ([]LocationGroup, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open location_group CSV: %w", err)
	}

	// Parse the CSV data into a slice of LocationGroup structs
	var locationGroups []LocationGroup
	for i := 0; i < len(df.Records); i++ {
		locationGroupID, err := df.GetString(i, "location_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'location_group_id' at row %d: %w", i, err)
		}

		locationGroupName, err := df.GetStringPtr(i, "location_group_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'location_group_name' at row %d: %w", i, err)
		}

		// Create LocationGroup struct and append to the list
		locationGroups = append(locationGroups, LocationGroup{
			LocationGroupId:   locationGroupID,
			LocationGroupName: locationGroupName,
		})
	}

	return locationGroups, nil
}

type LocationGroupGeom struct {
	LocationGroupId   string `gorm:"primaryKey"`
	LocationGroupName *string
	StopTimes         []StopTimesGeom         `gorm:"foreignKey:LocationGroupId;references:LocationGroupId"`
	LocationGroupStop []LocationGroupStopGeom `gorm:"foreignKey:LocationGroupId;references:LocationGroupId"`
}

func (LocationGroupGeom) TableName() string {
	return "location_group"
}

func ParseLocationGroupGeom(path string) ([]LocationGroupGeom, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open location_group CSV: %w", err)
	}

	// Parse the CSV data into a slice of LocationGroup structs
	var locationGroups []LocationGroupGeom
	for i := 0; i < len(df.Records); i++ {
		locationGroupID, err := df.GetString(i, "location_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'location_group_id' at row %d: %w", i, err)
		}

		locationGroupName, err := df.GetStringPtr(i, "location_group_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'location_group_name' at row %d: %w", i, err)
		}

		// Create LocationGroup struct and append to the list
		locationGroups = append(locationGroups, LocationGroupGeom{
			LocationGroupId:   locationGroupID,
			LocationGroupName: locationGroupName,
		})
	}

	return locationGroups, nil
}
