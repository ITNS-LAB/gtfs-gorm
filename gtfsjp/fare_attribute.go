package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareAttribute struct {
	FareId           string  `gorm:"primaryKey"`
	Price            float64 `gorm:"index;not null"`
	CurrencyType     string  `gorm:"primaryKey;default:JPY"`
	PaymentMethod    int     `gorm:"not null"`
	Transfers        *int    // Required field, but may be NULL if empty.
	TransferDuration *int
	//FareRule         FareRule `gorm:"foreignKey:FareId"`
}

func (FareAttribute) TableName() string {
	return "fare_attributes"
}

func ParseFareAttributes(path string) ([]FareAttribute, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open fare_attributes CSV: %w", err)
	}

	// データを解析して FareAttribute 構造体のスライスを作成
	var fareAttributes []FareAttribute
	for i := 0; i < len(df.Records); i++ {
		fareId, err := df.GetString(i, "fare_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_id' at row %d: %w", i, err)
		}

		price, err := df.GetFloat(i, "price")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'price' at row %d: %w", i, err)
		}

		currencyType, err := df.GetString(i, "currency_type")
		if err != nil || currencyType == "" {
			currencyType = "JPY" // デフォルト値
		}

		paymentMethod, err := df.GetInt(i, "payment_method")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'payment_method' at row %d: %w", i, err)
		}

		transfers, err := df.GetIntPtr(i, "transfers")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transfers' at row %d: %w", i, err)
		}

		transferDuration, err := df.GetIntPtr(i, "transfer_duration")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transfer_duration' at row %d: %w", i, err)
		}

		// FareAttribute 構造体を作成しリストに追加
		fareAttributes = append(fareAttributes, FareAttribute{
			FareId:           fareId,
			Price:            price,
			CurrencyType:     currencyType,
			PaymentMethod:    paymentMethod,
			Transfers:        transfers,
			TransferDuration: transferDuration,
		})
	}

	return fareAttributes, nil
}
