package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/csvutil"
	"gorm.io/datatypes"
)

type Frequency struct {
	TripId      string         `gorm:"primaryKey"`
	StartTime   datatypes.Time `gorm:"primaryKey"`
	EndTime     datatypes.Time `gorm:"primaryKey"`
	HeadwaySecs int            `gorm:"not null"`
	ExactTimes  *int           `gorm:"default:0"`
}

func (Frequency) TableName() string {
	return "frequencies"
}

func ParseFrequencies(path string) ([]Frequency, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open frequencies CSV: %w", err)
	}

	// データを解析して Frequency 構造体のスライスを作成
	var frequencies []Frequency
	for i := 0; i < len(df.Records); i++ {
		tripId, err := df.GetString(i, "trip_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_id' at row %d: %w", i, err)
		}

		startTime, err := df.GetTime(i, "start_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'start_time' at row %d: %w", i, err)
		}

		endTime, err := df.GetTime(i, "end_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'end_time' at row %d: %w", i, err)
		}

		headwaySecs, err := df.GetInt(i, "headway_secs")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'headway_secs' at row %d: %w", i, err)
		}

		exactTimes, err := df.GetIntPtr(i, "exact_times")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'exact_times' at row %d: %w", i, err)
		}

		// Frequency 構造体を作成しリストに追加
		frequencies = append(frequencies, Frequency{
			TripId:      tripId,
			StartTime:   startTime,
			EndTime:     endTime,
			HeadwaySecs: headwaySecs,
			ExactTimes:  exactTimes,
		})
	}

	return frequencies, nil
}
