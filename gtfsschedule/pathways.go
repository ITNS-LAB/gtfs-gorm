package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Pathway struct {
	PathwayID            string `gorm:"primaryKey"`
	FromStopID           string `gorm:"not null"`
	ToStopID             string `gorm:"not null"`
	PathwayMode          int    `gorm:"not null"`
	IsBidirectional      int    `gorm:"not null"`
	Length               *float64
	TraversalTime        *int
	StairCount           *int
	MaxSlope             *float64
	MinWidth             *float64
	SignpostedAs         *string
	ReversedSignpostedAs *string
}

func ParsePathway(path string) ([]Pathway, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open pathway CSV: %w", err)
	}

	// Parse the data and create a slice of Pathway structs
	var pathways []Pathway
	for i := 0; i < len(df.Records); i++ {
		pathwayID, err := df.GetString(i, "pathway_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'pathway_id' at row %d: %w", i, err)
		}

		fromStopID, err := df.GetString(i, "from_stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'from_stop_id' at row %d: %w", i, err)
		}

		toStopID, err := df.GetString(i, "to_stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'to_stop_id' at row %d: %w", i, err)
		}

		pathwayMode, err := df.GetInt(i, "pathway_mode")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'pathway_mode' at row %d: %w", i, err)
		}

		isBidirectional, err := df.GetInt(i, "is_bidirectional")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'is_bidirectional' at row %d: %w", i, err)
		}

		// Parse optional fields with pointers
		length, err := df.GetFloat64Ptr(i, "length")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'length' at row %d: %w", i, err)
		}

		traversalTime, err := df.GetIntPtr(i, "traversal_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'traversal_time' at row %d: %w", i, err)
		}

		stairCount, err := df.GetIntPtr(i, "stair_count")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stair_count' at row %d: %w", i, err)
		}

		maxSlope, err := df.GetFloat64Ptr(i, "max_slope")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'max_slope' at row %d: %w", i, err)
		}

		minWidth, err := df.GetFloat64Ptr(i, "min_width")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'min_width' at row %d: %w", i, err)
		}

		signpostedAs, err := df.GetStringPtr(i, "signposted_as")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'signposted_as' at row %d: %w", i, err)
		}

		reversedSignpostedAs, err := df.GetStringPtr(i, "reversed_signposted_as")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'reversed_signposted_as' at row %d: %w", i, err)
		}

		// Create the Pathway struct and append to the list
		pathways = append(pathways, Pathway{
			PathwayID:            pathwayID,
			FromStopID:           fromStopID,
			ToStopID:             toStopID,
			PathwayMode:          pathwayMode,
			IsBidirectional:      isBidirectional,
			Length:               length,
			TraversalTime:        traversalTime,
			StairCount:           stairCount,
			MaxSlope:             maxSlope,
			MinWidth:             minWidth,
			SignpostedAs:         signpostedAs,
			ReversedSignpostedAs: reversedSignpostedAs,
		})
	}

	return pathways, nil
}
