package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Transfer struct {
	Id              int    `gorm:"primaryKey;auto_increment"`
	FromStopId      string `gorm:"primaryKey"`
	ToStopId        string `gorm:"primaryKey"`
	TransferType    int    `gorm:"not null"`
	MinTransferTime *int
}

func (Transfer) TableName() string {
	return "transfers"
}

func ParseTransfers(path string) ([]Transfer, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open transfers CSV: %w", err)
	}

	// データを解析して Transfer 構造体のスライスを作成
	var transfers []Transfer
	for i := 0; i < len(df.Records); i++ {
		fromStopId, err := df.GetString(i, "from_stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'from_stop_id' at row %d: %w", i, err)
		}

		toStopId, err := df.GetString(i, "to_stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'to_stop_id' at row %d: %w", i, err)
		}

		transferType, err := df.GetInt(i, "transfer_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transfer_type' at row %d: %w", i, err)
		}

		minTransferTime, err := df.GetIntPtr(i, "min_transfer_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'min_transfer_time' at row %d: %w", i, err)
		}

		// Transfer 構造体を作成しリストに追加
		transfers = append(transfers, Transfer{
			FromStopId:      fromStopId,
			ToStopId:        toStopId,
			TransferType:    transferType,
			MinTransferTime: minTransferTime,
		})
	}

	return transfers, nil
}
