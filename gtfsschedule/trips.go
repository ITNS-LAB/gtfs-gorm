package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Trips struct {
	RouteId              string `gorm:"index;not null"`
	ServiceId            string `gorm:"index;not null"`
	TripId               string `gorm:"primaryKey"`
	TripHeadsign         *string
	TripShortName        *string
	DirectionId          *int    `gorm:"index"`
	BlockId              *string `gorm:"index"`
	ShapeId              string  `gorm:"index"`
	WheelchairAccessible *int
	BikesAllowed         *int
	StopTimes            []StopTimes   `gorm:"foreignKey:TripId;references:TripId"`
	Frequencies          []Frequencies `gorm:"foreignKey:TripId;references:TripId"`
	TransferFromTripID   []Transfer    `gorm:"foreignKey:FromTripId ;references:TripId"`
	TransferToTripID     []Transfer    `gorm:"foreignKey:ToTripId ;references:TripId"`
	Attribution          []Attribution `gorm:"foreignKey:TripId ;references:TripId"`
}

func ParseTrips(path string) ([]Trips, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open trips CSV: %w", err)
	}

	// データを解析してTrips構造体のスライスを作成
	var trips []Trips
	for i := 0; i < len(df.Records); i++ {
		routeId, err := df.GetString(i, "route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_id' at row %d: %w", i, err)
		}

		serviceId, err := df.GetString(i, "service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'service_id' at row %d: %w", i, err)
		}

		tripId, err := df.GetString(i, "trip_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_id' at row %d: %w", i, err)
		}

		tripHeadsign, err := df.GetStringPtr(i, "trip_headsign")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_headsign' at row %d: %w", i, err)
		}

		tripShortName, err := df.GetStringPtr(i, "trip_short_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_short_name' at row %d: %w", i, err)
		}

		directionId, err := df.GetIntPtr(i, "direction_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'direction_id' at row %d: %w", i, err)
		}

		blockId, err := df.GetStringPtr(i, "block_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'block_id' at row %d: %w", i, err)
		}

		shapeId, err := df.GetString(i, "shape_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'shape_id' at row %d: %w", i, err)
		}

		wheelchairAccessible, err := df.GetIntPtr(i, "wheelchair_accessible")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'wheelchair_accessible' at row %d: %w", i, err)
		}

		bikesAllowed, err := df.GetIntPtr(i, "bikes_allowed")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'bikes_allowed' at row %d: %w", i, err)
		}

		// Trips構造体を作成しリストに追加
		trips = append(trips, Trips{
			RouteId:              routeId,
			ServiceId:            serviceId,
			TripId:               tripId,
			TripHeadsign:         tripHeadsign,
			TripShortName:        tripShortName,
			DirectionId:          directionId,
			BlockId:              blockId,
			ShapeId:              shapeId,
			WheelchairAccessible: wheelchairAccessible,
			BikesAllowed:         bikesAllowed,
		})
	}

	return trips, nil
}
