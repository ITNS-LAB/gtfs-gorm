package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareTransferRule struct {
	FromLegGroupID    *string
	ToLegGroupID      *string
	TransferCount     *int
	DurationLimit     *int
	DurationLimitType *int
	FareTransferType  int `gorm:"not null"`
	FareProductID     *string
}

func ParseFareTransferRule(path string) ([]FareTransferRule, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open fare_transfer_rule CSV: %w", err)
	}

	// Parse the CSV data into a slice of FareTransferRule structs
	var fareTransferRules []FareTransferRule
	for i := 0; i < len(df.Records); i++ {
		fromLegGroupID, err := df.GetStringPtr(i, "from_leg_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'from_leg_group_id' at row %d: %w", i, err)
		}

		toLegGroupID, err := df.GetStringPtr(i, "to_leg_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'to_leg_group_id' at row %d: %w", i, err)
		}

		transferCount, err := df.GetIntPtr(i, "transfer_count")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transfer_count' at row %d: %w", i, err)
		}

		durationLimit, err := df.GetIntPtr(i, "duration_limit")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'duration_limit' at row %d: %w", i, err)
		}

		durationLimitType, err := df.GetIntPtr(i, "duration_limit_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'duration_limit_type' at row %d: %w", i, err)
		}

		fareTransferType, err := df.GetInt(i, "fare_transfer_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_transfer_type' at row %d: %w", i, err)
		}

		fareProductID, err := df.GetStringPtr(i, "fare_product_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_product_id' at row %d: %w", i, err)
		}

		// Create FareTransferRule struct and append to the list
		fareTransferRules = append(fareTransferRules, FareTransferRule{
			FromLegGroupID:    fromLegGroupID,
			ToLegGroupID:      toLegGroupID,
			TransferCount:     transferCount,
			DurationLimit:     durationLimit,
			DurationLimitType: durationLimitType,
			FareTransferType:  fareTransferType,
			FareProductID:     fareProductID,
		})
	}

	return fareTransferRules, nil
}

type FareTransferRuleGeom struct {
	FromLegGroupID    *string
	ToLegGroupID      *string
	TransferCount     *int
	DurationLimit     *int
	DurationLimitType *int
	FareTransferType  int `gorm:"not null"`
	FareProductID     *string
}

func ParseFareTransferRuleGeom(path string) ([]FareTransferRuleGeom, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open fare_transfer_rule CSV: %w", err)
	}

	// Parse the CSV data into a slice of FareTransferRule structs
	var fareTransferRules []FareTransferRuleGeom
	for i := 0; i < len(df.Records); i++ {
		fromLegGroupID, err := df.GetStringPtr(i, "from_leg_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'from_leg_group_id' at row %d: %w", i, err)
		}

		toLegGroupID, err := df.GetStringPtr(i, "to_leg_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'to_leg_group_id' at row %d: %w", i, err)
		}

		transferCount, err := df.GetIntPtr(i, "transfer_count")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transfer_count' at row %d: %w", i, err)
		}

		durationLimit, err := df.GetIntPtr(i, "duration_limit")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'duration_limit' at row %d: %w", i, err)
		}

		durationLimitType, err := df.GetIntPtr(i, "duration_limit_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'duration_limit_type' at row %d: %w", i, err)
		}

		fareTransferType, err := df.GetInt(i, "fare_transfer_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_transfer_type' at row %d: %w", i, err)
		}

		fareProductID, err := df.GetStringPtr(i, "fare_product_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_product_id' at row %d: %w", i, err)
		}

		// Create FareTransferRule struct and append to the list
		fareTransferRules = append(fareTransferRules, FareTransferRuleGeom{
			FromLegGroupID:    fromLegGroupID,
			ToLegGroupID:      toLegGroupID,
			TransferCount:     transferCount,
			DurationLimit:     durationLimit,
			DurationLimitType: durationLimitType,
			FareTransferType:  fareTransferType,
			FareProductID:     fareProductID,
		})
	}

	return fareTransferRules, nil
}
