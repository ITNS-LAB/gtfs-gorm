package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Trip struct {
	RouteId              string `gorm:"index;not null"`
	ServiceId            string `gorm:"index;not null"`
	TripId               string `gorm:"primaryKey"`
	TripHeadsign         *string
	TripShortName        *string
	DirectionId          *int    `gorm:"index"`
	BlockId              *string `gorm:"index"`
	ShapeId              *string `gorm:"index"`
	WheelchairAccessible *int    `gorm:"default:0"`
	BikesAllowed         *int    `gorm:"default:0"`
	JpTripDesc           *string
	JpTripDescSymbol     *string
	JpOfficeId           *string
	JpPatternId          *string
	StopTimes            []StopTime  `gorm:"foreignKey:TripId;references:TripId"`
	Frequencies          []Frequency `gorm:"foreignKey:TripId;references:TripId"`
}

func (Trip) TableName() string {
	return "trips"
}

func ParseTrips(path string) ([]Trip, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open trips CSV: %w", err)
	}

	// データを解析して Trip 構造体のスライスを作成
	var trips []Trip
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

		shapeId, err := df.GetStringPtr(i, "shape_id")
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

		jpTripDesc, err := df.GetStringPtr(i, "jp_trip_desc")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_trip_desc' at row %d: %w", i, err)
		}

		jpTripDescSymbol, err := df.GetStringPtr(i, "jp_trip_desc_symbol")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_trip_desc_symbol' at row %d: %w", i, err)
		}

		jpOfficeId, err := df.GetStringPtr(i, "jp_office_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_office_id' at row %d: %w", i, err)
		}

		jpPatternId, err := df.GetStringPtr(i, "jp_pattern_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_pattern_id' at row %d: %w", i, err)
		}

		// Trip 構造体を作成しリストに追加
		trips = append(trips, Trip{
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
			JpTripDesc:           jpTripDesc,
			JpTripDescSymbol:     jpTripDescSymbol,
			JpOfficeId:           jpOfficeId,
			JpPatternId:          jpPatternId,
		})
	}

	return trips, nil
}

type TripGeom struct {
	RouteId              string `gorm:"index;not null"`
	ServiceId            string `gorm:"index;not null"`
	TripId               string `gorm:"primaryKey"`
	TripHeadsign         *string
	TripShortName        *string
	DirectionId          *int    `gorm:"index"`
	BlockId              *string `gorm:"index"`
	ShapeId              *string `gorm:"index"`
	WheelchairAccessible *int    `gorm:"default:0"`
	BikesAllowed         *int    `gorm:"default:0"`
	JpTripDesc           *string
	JpTripDescSymbol     *string
	JpOfficeId           *string
	JpPatternId          *string
	Geom                 gormdatatypes.Geometry
	StopTimes            []StopTime  `gorm:"foreignKey:TripId;references:TripId"`
	Frequencies          []Frequency `gorm:"foreignKey:TripId;references:TripId"`
}

func (TripGeom) TableName() string {
	return "trips_geom"
}

func ParseTripsGeom(path string) ([]TripGeom, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open trips CSV: %w", err)
	}

	// データを解析して Trip 構造体のスライスを作成
	var tripsGeom []TripGeom
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

		shapeId, err := df.GetStringPtr(i, "shape_id")
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

		jpTripDesc, err := df.GetStringPtr(i, "jp_trip_desc")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_trip_desc' at row %d: %w", i, err)
		}

		jpTripDescSymbol, err := df.GetStringPtr(i, "jp_trip_desc_symbol")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_trip_desc_symbol' at row %d: %w", i, err)
		}

		jpOfficeId, err := df.GetStringPtr(i, "jp_office_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_office_id' at row %d: %w", i, err)
		}

		jpPatternId, err := df.GetStringPtr(i, "jp_pattern_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_pattern_id' at row %d: %w", i, err)
		}

		// Trip 構造体を作成しリストに追加
		tripsGeom = append(tripsGeom, TripGeom{
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
			JpTripDesc:           jpTripDesc,
			JpTripDescSymbol:     jpTripDescSymbol,
			JpOfficeId:           jpOfficeId,
			JpPatternId:          jpPatternId,
		})
	}

	return tripsGeom, nil
}
