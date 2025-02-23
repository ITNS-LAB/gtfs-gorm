package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareProduct struct {
	FareProductID   string `gorm:"primaryKey"`
	FareProductName *string
	FareMediaID     *string
	Amount          float64 `gorm:"not null"`
	Currency        string  `gorm:"not null"`
	//FareLeg          []FareLeg          `gorm:"foreignKey:FareProductID;references:FareProductID"`
	FareTransferRule []FareTransferRule `gorm:"foreignKey:FareProductID;references:FareProductID"`
}

func ParseFareProduct(path string) ([]FareProduct, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open FareProduct CSV: %w", err)
	}

	// データを解析してFareProduct構造体のスライスを作成
	var fareProducts []FareProduct
	for i := 0; i < len(df.Records); i++ {
		fareProductID, err := df.GetString(i, "fare_product_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_product_id' at row %d: %w", i, err)
		}

		fareProductName, err := df.GetStringPtr(i, "fare_product_name")
		if err != nil {
			// Optional field, so it's okay if this is nil
		}

		fareMediaID, err := df.GetStringPtr(i, "fare_media_id")
		if err != nil {
			// Optional field, so it's okay if this is nil
		}

		amount, err := df.GetFloat(i, "amount")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'amount' at row %d: %w", i, err)
		}

		currency, err := df.GetString(i, "currency")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'currency' at row %d: %w", i, err)
		}

		// FareProduct 構造体を作成しリストに追加
		fareProducts = append(fareProducts, FareProduct{
			FareProductID:   fareProductID,
			FareProductName: fareProductName,
			FareMediaID:     fareMediaID,
			Amount:          amount,
			Currency:        currency,
		})
	}

	return fareProducts, nil
}

type FareProductGeom struct {
	FareProductID   string `gorm:"primaryKey"`
	FareProductName *string
	FareMediaID     *string
	Amount          float64 `gorm:"not null"`
	Currency        string  `gorm:"not null"`
	//FareLeg          []FareLeg          `gorm:"foreignKey:FareProductID;references:FareProductID"`
	FareTransferRule []FareTransferRuleGeom `gorm:"foreignKey:FareProductID;references:FareProductID"`
}

func ParseFareProductGeom(path string) ([]FareProductGeom, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open FareProduct CSV: %w", err)
	}

	// データを解析してFareProduct構造体のスライスを作成
	var fareProducts []FareProductGeom
	for i := 0; i < len(df.Records); i++ {
		fareProductID, err := df.GetString(i, "fare_product_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_product_id' at row %d: %w", i, err)
		}

		fareProductName, err := df.GetStringPtr(i, "fare_product_name")
		if err != nil {
			// Optional field, so it's okay if this is nil
		}

		fareMediaID, err := df.GetStringPtr(i, "fare_media_id")
		if err != nil {
			// Optional field, so it's okay if this is nil
		}

		amount, err := df.GetFloat(i, "amount")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'amount' at row %d: %w", i, err)
		}

		currency, err := df.GetString(i, "currency")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'currency' at row %d: %w", i, err)
		}

		// FareProduct 構造体を作成しリストに追加
		fareProducts = append(fareProducts, FareProductGeom{
			FareProductID:   fareProductID,
			FareProductName: fareProductName,
			FareMediaID:     fareMediaID,
			Amount:          amount,
			Currency:        currency,
		})
	}

	return fareProducts, nil
}
