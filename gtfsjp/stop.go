package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"github.com/paulmach/orb"
)

type Stop struct {
	StopId             string `gorm:"primaryKey"`
	StopCode           *string
	StopName           string `gorm:"not null"`
	StopDesc           *string
	StopLat            float64 `gorm:"not null"`
	StopLon            float64 `gorm:"not null"`
	ZoneId             *string `gorm:"unique"`
	StopUrl            *string
	LocationType       *int `gorm:"default:0"`
	ParentStation      *string
	StopTimezone       *string
	WheelchairBoarding *int `gorm:"default:0"`
	LevelId            *string
	PlatformCode       *string
	StopTimes          []StopTime `gorm:"foreignKey:StopId;references:StopId"`
}

func (Stop) TableName() string {
	return "stops"
}

func ParseStops(path string) ([]Stop, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stops CSV: %w", err)
	}

	// データを解析して Stop 構造体のスライスを作成
	var stops []Stop
	for i := 0; i < len(df.Records); i++ {
		stopId, err := df.GetString(i, "stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_id' at row %d: %w", i, err)
		}

		stopCode, err := df.GetStringPtr(i, "stop_code")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_code' at row %d: %w", i, err)
		}

		stopName, err := df.GetString(i, "stop_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_name' at row %d: %w", i, err)
		}

		stopDesc, err := df.GetStringPtr(i, "stop_desc")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_desc' at row %d: %w", i, err)
		}

		stopLat, err := df.GetFloat(i, "stop_lat")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_lat' at row %d: %w", i, err)
		}

		stopLon, err := df.GetFloat(i, "stop_lon")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_lon' at row %d: %w", i, err)
		}

		zoneId, err := df.GetStringPtr(i, "zone_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'zone_id' at row %d: %w", i, err)
		}

		stopUrl, err := df.GetStringPtr(i, "stop_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_url' at row %d: %w", i, err)
		}

		locationType, err := df.GetIntPtr(i, "location_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'location_type' at row %d: %w", i, err)
		}

		parentStation, err := df.GetStringPtr(i, "parent_station")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'parent_station' at row %d: %w", i, err)
		}

		stopTimezone, err := df.GetStringPtr(i, "stop_timezone")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_timezone' at row %d: %w", i, err)
		}

		wheelchairBoarding, err := df.GetIntPtr(i, "wheelchair_boarding")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'wheelchair_boarding' at row %d: %w", i, err)
		}

		levelId, err := df.GetStringPtr(i, "level_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'level_id' at row %d: %w", i, err)
		}

		platformCode, err := df.GetStringPtr(i, "platform_code")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'platform_code' at row %d: %w", i, err)
		}

		// Stop 構造体を作成しリストに追加
		stops = append(stops, Stop{
			StopId:             stopId,
			StopCode:           stopCode,
			StopName:           stopName,
			StopDesc:           stopDesc,
			StopLat:            stopLat,
			StopLon:            stopLon,
			ZoneId:             zoneId,
			StopUrl:            stopUrl,
			LocationType:       locationType,
			ParentStation:      parentStation,
			StopTimezone:       stopTimezone,
			WheelchairBoarding: wheelchairBoarding,
			LevelId:            levelId,
			PlatformCode:       platformCode,
		})
	}

	return stops, nil
}

type StopGeom struct {
	StopId             string `gorm:"primaryKey"`
	StopCode           *string
	StopName           string `gorm:"not null"`
	StopDesc           *string
	StopLat            float64 `gorm:"not null"`
	StopLon            float64 `gorm:"not null"`
	ZoneId             *string `gorm:"unique"`
	StopUrl            *string
	LocationType       *int `gorm:"default:0"`
	ParentStation      *string
	StopTimezone       *string
	WheelchairBoarding *int `gorm:"default:0"`
	LevelId            *string
	PlatformCode       *string
	Geom               gormdatatypes.Geometry `gorm:"index"`
	StopTimes          []StopTime             `gorm:"foreignKey:StopId;references:StopId"`
}

func (StopGeom) TableName() string {
	return "stops_geom"
}

func ParseStopsGeom(path string) ([]StopGeom, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stops CSV: %w", err)
	}

	// データを解析して Stop 構造体のスライスを作成
	var stopsGeom []StopGeom
	for i := 0; i < len(df.Records); i++ {
		stopId, err := df.GetString(i, "stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_id' at row %d: %w", i, err)
		}

		stopCode, err := df.GetStringPtr(i, "stop_code")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_code' at row %d: %w", i, err)
		}

		stopName, err := df.GetString(i, "stop_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_name' at row %d: %w", i, err)
		}

		stopDesc, err := df.GetStringPtr(i, "stop_desc")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_desc' at row %d: %w", i, err)
		}

		stopLat, err := df.GetFloat(i, "stop_lat")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_lat' at row %d: %w", i, err)
		}

		stopLon, err := df.GetFloat(i, "stop_lon")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_lon' at row %d: %w", i, err)
		}

		zoneId, err := df.GetStringPtr(i, "zone_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'zone_id' at row %d: %w", i, err)
		}

		stopUrl, err := df.GetStringPtr(i, "stop_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_url' at row %d: %w", i, err)
		}

		locationType, err := df.GetIntPtr(i, "location_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'location_type' at row %d: %w", i, err)
		}

		parentStation, err := df.GetStringPtr(i, "parent_station")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'parent_station' at row %d: %w", i, err)
		}

		stopTimezone, err := df.GetStringPtr(i, "stop_timezone")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_timezone' at row %d: %w", i, err)
		}

		wheelchairBoarding, err := df.GetIntPtr(i, "wheelchair_boarding")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'wheelchair_boarding' at row %d: %w", i, err)
		}

		levelId, err := df.GetStringPtr(i, "level_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'level_id' at row %d: %w", i, err)
		}

		platformCode, err := df.GetStringPtr(i, "platform_code")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'platform_code' at row %d: %w", i, err)
		}

		geom := gormdatatypes.Geometry{
			Geom: orb.Point{stopLon, stopLat},
			Srid: 4326,
		}

		// Stop 構造体を作成しリストに追加
		stopsGeom = append(stopsGeom, StopGeom{
			StopId:             stopId,
			StopCode:           stopCode,
			StopName:           stopName,
			StopDesc:           stopDesc,
			StopLat:            stopLat,
			StopLon:            stopLon,
			ZoneId:             zoneId,
			StopUrl:            stopUrl,
			LocationType:       locationType,
			ParentStation:      parentStation,
			StopTimezone:       stopTimezone,
			WheelchairBoarding: wheelchairBoarding,
			LevelId:            levelId,
			PlatformCode:       platformCode,
			Geom:               geom,
		})
	}

	return stopsGeom, nil
}
