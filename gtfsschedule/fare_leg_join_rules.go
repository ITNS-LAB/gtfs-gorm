package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareLegJoinRules struct {
	FromNetworkId string `gorm:"primaryKey"`
	ToNetworkId   string `gorm:"not null"`
	FromStopId    *string
	ToStopId      *string
}

func (FareLegJoinRules) TableName() string {
	return "fareLegJoinRules"
}

func ParseFareLegJoinRules(path string) ([]FareLegJoinRules, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open FareLegJoinRules CSV: %w", err)
	}

	// データを解析してFareLegJoinRules構造体のスライスを作成
	var fareLegJoinRules []FareLegJoinRules
	for i := 0; i < len(df.Records); i++ {
		fromNetworkID, err := df.GetString(i, "from_network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'from_network_id' at row %d: %w", i, err)
		}

		toNetworkID, err := df.GetString(i, "to_network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'to_network_id' at row %d: %w", i, err)
		}

		fromStopID, err := df.GetStringPtr(i, "from_stop_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		toStopID, err := df.GetStringPtr(i, "to_stop_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		// FareLegJoinRules 構造体を作成しリストに追加
		fareLegJoinRules = append(fareLegJoinRules, FareLegJoinRules{
			FromNetworkId: fromNetworkID,
			ToNetworkId:   toNetworkID,
			FromStopId:    fromStopID,
			ToStopId:      toStopID,
		})
	}

	return fareLegJoinRules, nil
}

type FareLegJoinRulesGeom struct {
	FromNetworkId string `gorm:"primaryKey"`
	ToNetworkId   string `gorm:"not null"`
	FromStopId    *string
	ToStopId      *string
}

func (FareLegJoinRulesGeom) TableName() string {
	return "fareLegJoinRules"
}

func ParseFareLegJoinRulesGeom(path string) ([]FareLegJoinRulesGeom, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open FareLegJoinRules CSV: %w", err)
	}

	// データを解析してFareLegJoinRules構造体のスライスを作成
	var fareLegJoinRules []FareLegJoinRulesGeom
	for i := 0; i < len(df.Records); i++ {
		fromNetworkID, err := df.GetString(i, "from_network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'from_network_id' at row %d: %w", i, err)
		}

		toNetworkID, err := df.GetString(i, "to_network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'to_network_id' at row %d: %w", i, err)
		}

		fromStopID, err := df.GetStringPtr(i, "from_stop_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		toStopID, err := df.GetStringPtr(i, "to_stop_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		// FareLegJoinRules 構造体を作成しリストに追加
		fareLegJoinRules = append(fareLegJoinRules, FareLegJoinRulesGeom{
			FromNetworkId: fromNetworkID,
			ToNetworkId:   toNetworkID,
			FromStopId:    fromStopID,
			ToStopId:      toStopID,
		})
	}

	return fareLegJoinRules, nil
}
