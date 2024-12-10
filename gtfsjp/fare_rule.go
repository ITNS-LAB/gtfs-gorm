package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareRule struct {
	Id            int    `gorm:"primaryKey;auto_increment"`
	FareId        string `gorm:"index;not null"`
	RouteId       *string
	OriginId      *string
	DestinationId *string
	ContainsId    *string // Not used
}

func (FareRule) TableName() string {
	return "fare_rules"
}

func ParseFareRules(path string) ([]FareRule, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open fare_rules CSV: %w", err)
	}

	// データを解析して FareRule 構造体のスライスを作成
	var fareRules []FareRule
	for i := 0; i < len(df.Records); i++ {
		fareId, err := df.GetString(i, "fare_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_id' at row %d: %w", i, err)
		}

		routeId, err := df.GetStringPtr(i, "route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_id' at row %d: %w", i, err)
		}

		originId, err := df.GetStringPtr(i, "origin_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'origin_id' at row %d: %w", i, err)
		}

		destinationId, err := df.GetStringPtr(i, "destination_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'destination_id' at row %d: %w", i, err)
		}

		containsId, err := df.GetStringPtr(i, "contains_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'contains_id' at row %d: %w", i, err)
		}

		// FareRule 構造体を作成しリストに追加
		fareRules = append(fareRules, FareRule{
			FareId:        fareId,
			RouteId:       routeId,
			OriginId:      originId,
			DestinationId: destinationId,
			ContainsId:    containsId,
		})
	}

	return fareRules, nil
}
