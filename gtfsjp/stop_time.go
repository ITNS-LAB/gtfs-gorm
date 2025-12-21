package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type StopTimeTemp struct {
	TripId        string `gorm:"primaryKey"`
	ArrivalTime   string `gorm:"index;not null"`
	DepartureTime string `gorm:"index;not null"`
	/*
		ArrivalTime       datatypes.Time `gorm:"index;not null"`
		DepartureTime     datatypes.Time `gorm:"index;not null"`
	*/
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

func ParseStopTimes(path string) ([]StopTime, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stop_times CSV: %w", err)
	}

	// データを解析して StopTime 構造体のスライスを作成
	var stopTimes []StopTimeTemp
	for i := 0; i < len(df.Records); i++ {
		tripId, err := df.GetString(i, "trip_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_id' at row %d: %w", i, err)
		}

		arrivalTime, err := df.GetString(i, "arrival_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'arrival_time' at row %d: %w", i, err)
		}

		departureTime, err := df.GetString(i, "departure_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'departure_time' at row %d: %w", i, err)
		}

		// stop_timeのhour out of rangeへの一時的な対処
		/*
			arrivalTime, err := df.GetTime(i, "arrival_time")
			if err != nil {
				return nil, fmt.Errorf("failed to get 'arrival_time' at row %d: %w", i, err)
			}

			departureTime, err := df.GetTime(i, "departure_time")
			if err != nil {
				return nil, fmt.Errorf("failed to get 'departure_time' at row %d: %w", i, err)
			}
		*/

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
		stopTimes = append(stopTimes, StopTimeTemp{
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

	// 構造体から路線がドリーミントの部分だけを取り除く処理をして返す。
	var stop_time_filtered []StopTime
	for _, stopTime := range stopTimes {
		if stopTime.TripId == "土日祝_17時55分_系統901" ||
			stopTime.TripId == "平日_17時55分_系統901" ||
			stopTime.TripId == "土日祝_19時20分_系統905" ||
			stopTime.TripId == "平日_19時20分_系統905" {
			continue
		}
		var h, m, s int
		_, err := fmt.Sscanf(stopTime.ArrivalTime, "%d:%d:%d", &h, &m, &s)
		if err != nil {
			return nil, err
		}
		arrival := datatypes.NewTime(h, m, s, 0)
		_, err = fmt.Sscanf(stopTime.DepartureTime, "%d:%d:%d", &h, &m, &s)
		if err != nil {
			return nil, err
		}
		departure := datatypes.NewTime(h, m, s, 0)

		stop_time_filtered = append(stop_time_filtered, StopTime{
			TripId:            stopTime.TripId,
			ArrivalTime:       arrival,
			DepartureTime:     departure,
			StopId:            stopTime.StopId,
			StopSequence:      stopTime.StopSequence,
			StopHeadsign:      stopTime.StopHeadsign,
			PickupType:        stopTime.PickupType,
			DropOffType:       stopTime.DropOffType,
			ContinuousPickup:  stopTime.ContinuousPickup,
			ContinuousDropOff: stopTime.ContinuousDropOff,
			ShapeDistTraveled: stopTime.ShapeDistTraveled,
			Timepoint:         stopTime.Timepoint,
		})
	}

	return stop_time_filtered, nil
}

type StopTimeGeom struct {
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

func (StopTimeGeom) TableName() string {
	return "stop_times"
}

func ParseStopTimesGeom(path string) ([]StopTimeGeom, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stop_times CSV: %w", err)
	}

	// データを解析して StopTime 構造体のスライスを作成
	var stopTimes []StopTimeGeom
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
		stopTimes = append(stopTimes, StopTimeGeom{
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
