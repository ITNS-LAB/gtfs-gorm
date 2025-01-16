package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareRules struct {
	FareID        int `gorm:"primary_key"`
	RouteID       *string
	OriginID      *string
	DestinationID *string
	ContainsId    *string
}

func ParseFareRules(path string) ([]FareRules, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open fare_rules CSV: %w", err)
	}

	// データを解析して FareRules 構造体のスライスを作成
	var fareRules []FareRules
	for i := 0; i < len(df.Records); i++ {
		fareID, err := df.GetInt(i, "fare_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_id' at row %d: %w", i, err)
		}

		routeID, err := df.GetStringPtr(i, "route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_id' at row %d: %w", i, err)
		}

		originID, err := df.GetStringPtr(i, "origin_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'origin_id' at row %d: %w", i, err)
		}

		destinationID, err := df.GetStringPtr(i, "destination_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'destination_id' at row %d: %w", i, err)
		}

		containsId, err := df.GetStringPtr(i, "contains_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'contains_id' at row %d: %w", i, err)
		}

		// FareRules 構造体を作成しリストに追加
		fareRules = append(fareRules, FareRules{
			FareID:        fareID,
			RouteID:       routeID,
			OriginID:      originID,
			DestinationID: destinationID,
			ContainsId:    containsId,
		})
	}

	return fareRules, nil
}
