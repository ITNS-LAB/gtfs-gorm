package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Transfer struct {
	FromStopID      string
	ToStopID        string
	FromRouteID     *string
	ToRouteID       *string
	FromTripID      *string
	ToTripID        *string
	TransferType    int `gorm:"not null"` // 接続タイプを示します (0, 1, 2, 3, 4, 5)
	MinTransferTime *int
}

func ParseTransfer(path string) ([]Transfer, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open transfer CSV: %w", err)
	}

	// データを解析してTransfer構造体のスライスを作成
	var transfers []Transfer
	for i := 0; i < len(df.Records); i++ {
		fromStopID, err := df.GetString(i, "from_stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'from_stop_id' at row %d: %w", i, err)
		}

		toStopID, err := df.GetString(i, "to_stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'to_stop_id' at row %d: %w", i, err)
		}

		fromRouteID, err := df.GetStringPtr(i, "from_route_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		toRouteID, err := df.GetStringPtr(i, "to_route_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		fromTripID, err := df.GetStringPtr(i, "from_trip_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		toTripID, err := df.GetStringPtr(i, "to_trip_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		transferTypeStr, err := df.GetString(i, "transfer_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transfer_type' at row %d: %w", i, err)
		}
		transferType := 0
		// 文字列をintに変換
		_, err = fmt.Sscanf(transferTypeStr, "%d", &transferType)
		if err != nil {
			return nil, fmt.Errorf("failed to convert 'transfer_type' to int at row %d: %w", i, err)
		}

		minTransferTime, err := df.GetIntPtr(i, "min_transfer_time")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		// Transfer 構造体を作成しリストに追加
		transfers = append(transfers, Transfer{
			FromStopID:      fromStopID,
			ToStopID:        toStopID,
			FromRouteID:     fromRouteID,
			ToRouteID:       toRouteID,
			FromTripID:      fromTripID,
			ToTripID:        toTripID,
			TransferType:    transferType,
			MinTransferTime: minTransferTime,
		})
	}

	return transfers, nil
}
