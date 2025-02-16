package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareAttributes struct {
	FareId           string  `gorm:"primaryKey"`
	Price            float64 `gorm:"not null"`
	CurrencyType     string  `gorm:"not null"`
	PaymentMethod    int     `gorm:"not null"`
	Transfers        int     `gorm:"not null"`
	AgencyId         *string
	TransferDuration *int
	FareRules        []FareRules `gorm:"foreignKey:FareId;references:FareId"`
}

func ParseFareAttributes(path string) ([]FareAttributes, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open fare_attributes CSV: %w", err)
	}

	// データを解析して FareAttributes 構造体のスライスを作成
	var fareAttributes []FareAttributes
	for i := 0; i < len(df.Records); i++ {
		fareID, err := df.GetString(i, "fare_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_id' at row %d: %w", i, err)
		}

		price, err := df.GetFloat(i, "price")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'price' at row %d: %w", i, err)
		}

		currencyType, err := df.GetString(i, "currency_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'currency_type' at row %d: %w", i, err)
		}

		paymentMethod, err := df.GetInt(i, "payment_method")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'payment_method' at row %d: %w", i, err)
		}

		transfers, err := df.GetInt(i, "transfers")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transfers' at row %d: %w", i, err)
		}

		agencyID, err := df.GetStringPtr(i, "agency_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_id' at row %d: %w", i, err)
		}

		transferDuration, err := df.GetIntPtr(i, "transferDuration")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'transferDuration' at row %d: %w", i, err)
		}

		// FareAttributes 構造体を作成しリストに追加
		fareAttributes = append(fareAttributes, FareAttributes{
			FareId:           fareID,
			Price:            price,
			CurrencyType:     currencyType,
			PaymentMethod:    paymentMethod,
			Transfers:        transfers,
			AgencyId:         agencyID,
			TransferDuration: transferDuration,
		})
	}

	return fareAttributes, nil
}
