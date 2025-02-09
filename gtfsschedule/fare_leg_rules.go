package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareLegRules struct {
	LegGroupID                     *string
	NetworkID                      *string
	FromAreaID                     *string
	ToAreaID                       *string
	FromTimeframeGroupID           *string
	ToTimeframeGroupID             *string
	FareProductID                  string `gorm:"not null"`
	RulePriority                   *int
	FareTransferRuleFromLegGroupID []FareTransferRule `gorm:"foreignKey:LegGroupID;references:FromLegGroupID "`
	FareTransferRuleToLegGroupID   []FareTransferRule `gorm:"foreignKey:LegGroupID;references:ToLegGroupID "`
}

func ParseFareLeg(path string) ([]FareLegRules, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open FareLeg CSV: %w", err)
	}

	// データを解析してFareLeg構造体のスライスを作成
	var fareLegs []FareLegRules
	for i := 0; i < len(df.Records); i++ {
		legGroupID, err := df.GetStringPtr(i, "leg_group_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		networkID, err := df.GetStringPtr(i, "network_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		fromAreaID, err := df.GetStringPtr(i, "from_area_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		toAreaID, err := df.GetStringPtr(i, "to_area_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		fromTimeframeGroupID, err := df.GetStringPtr(i, "from_timeframe_group_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		toTimeframeGroupID, err := df.GetStringPtr(i, "to_timeframe_group_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		fareProductID, err := df.GetString(i, "fare_product_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_product_id' at row %d: %w", i, err)
		}

		rulePriority, err := df.GetIntPtr(i, "rule_priority")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		// FareLeg 構造体を作成しリストに追加
		fareLegs = append(fareLegs, FareLegRules{
			LegGroupID:           legGroupID,
			NetworkID:            networkID,
			FromAreaID:           fromAreaID,
			ToAreaID:             toAreaID,
			FromTimeframeGroupID: fromTimeframeGroupID,
			ToTimeframeGroupID:   toTimeframeGroupID,
			FareProductID:        fareProductID,
			RulePriority:         rulePriority,
		})
	}

	return fareLegs, nil
}
