package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type StopTime struct {
	TripId            string         `gorm:"primaryKey"`
	ArrivalTime       datatypes.Time `gorm:"index;not null"`
	DepartureTime     datatypes.Time `gorm:"index;not null"`
	StopId            string
	StopSequence      int `gorm:"primaryKey"`
	StopHeadsign      *string
	PickupType        *int `gorm:"default:0"`
	DropOffType       *int `gorm:"default:0"`
	ContinuousPickup  *int `gorm:"default:1"`
	ContinuousDropOff *int `gorm:"default:1"`
	ShapeDistTraveled *float64
	Timepoint         *int `gorm:"default:1"` // Not used in Japan (GTFS-JP)
}

func (StopTime) TableName() string {
	return "stop_times"
}

func ParseStopTimes(path string) ([]StopTime, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stop_times CSV: %w", err)
	}

	// データを解析して StopTime 構造体のスライスを作成
	var stopTimes []StopTime
	for i := 0; i < len(df.Records); i++ {
		tripId, err := df.GetString(i, "trip_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_id' at row %d: %w", i, err)
		}

		arrivalTime, err := df.GetTime(i, "arrival_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'arrival_time' at row %d: %w", i, err)
		}

		departureTime, err := df.GetTime(i, "departure_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'departure_time' at row %d: %w", i, err)
		}

		stopId, err := df.GetString(i, "stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_id' at row %d: %w", i, err)
		}

		stopSequence, err := df.GetInt(i, "stop_sequence")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_sequence' at row %d: %w", i, err)
		}

		stopHeadsign, err := df.GetStringPtr(i, "stop_headsign")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_headsign' at row %d: %w", i, err)
		}

		pickupType, err := df.GetIntPtr(i, "pickup_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'pickup_type' at row %d: %w", i, err)
		}

		dropOffType, err := df.GetIntPtr(i, "drop_off_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'drop_off_type' at row %d: %w", i, err)
		}

		continuousPickup, err := df.GetIntPtr(i, "continuous_pickup")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'continuous_pickup' at row %d: %w", i, err)
		}

		continuousDropOff, err := df.GetIntPtr(i, "continuous_drop_off")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'continuous_drop_off' at row %d: %w", i, err)
		}

		shapeDistTraveled, err := df.GetFloatPtr(i, "shape_dist_traveled")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'shape_dist_traveled' at row %d: %w", i, err)
		}

		timepoint, err := df.GetIntPtr(i, "timepoint")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'timepoint' at row %d: %w", i, err)
		}

		// StopTime 構造体を作成しリストに追加
		stopTimes = append(stopTimes, StopTime{
			TripId:            tripId,
			ArrivalTime:       arrivalTime,
			DepartureTime:     departureTime,
			StopId:            stopId,
			StopSequence:      stopSequence,
			StopHeadsign:      stopHeadsign,
			PickupType:        pickupType,
			DropOffType:       dropOffType,
			ContinuousPickup:  continuousPickup,
			ContinuousDropOff: continuousDropOff,
			ShapeDistTraveled: shapeDistTraveled,
			Timepoint:         timepoint,
		})
	}

	return stopTimes, nil
}
