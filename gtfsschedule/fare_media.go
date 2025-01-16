package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type FareMedia struct {
	FareMediaID   string `gorm:"primary_key"`
	FareMediaName *string
	FareMediaType int           `gorm:"not null"`
	FareProduct   []FareProduct `gorm:"foreignKey:FareMediaID;references:FareMediaID "`
}

func ParseFareMedia(path string) ([]FareMedia, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open FareMedia CSV: %w", err)
	}

	// データを解析してFareMedia構造体のスライスを作成
	var fareMedias []FareMedia
	for i := 0; i < len(df.Records); i++ {
		fareMediaID, err := df.GetString(i, "fare_media_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_media_id' at row %d: %w", i, err)
		}

		fareMediaName, err := df.GetStringPtr(i, "fare_media_name")
		if err != nil {
			// Optional field, so it's okay if this is nil
		}

		fareMediaType, err := df.GetInt(i, "fare_media_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'fare_media_type' at row %d: %w", i, err)
		}

		// FareMedia 構造体を作成しリストに追加
		fareMedias = append(fareMedias, FareMedia{
			FareMediaID:   fareMediaID,
			FareMediaName: fareMediaName,
			FareMediaType: fareMediaType,
		})
	}

	return fareMedias, nil
}
