package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Stop struct {
	StopId                     string `gorm:"primary_key"`
	StopCode                   *string
	StopName                   string `gorm:"not null"`
	TtsStopName                *string
	StopDesc                   *string
	StopLat                    float64 `gorm:"not null"`
	StopLon                    float64 `gorm:"not null"`
	ZoneId                     string  `gorm:"index;unique"`
	StopUrl                    *string
	LocationType               *int
	ParentStation              *int
	StopTimezone               *string
	WheelchairBoarding         *int
	LevelId                    *string
	PlatformCode               *string
	StopTimes                  []StopTimes         `gorm:"foreignKey:StopId;references:StopId"`
	FareRulesOriginID          FareRules           `gorm:"foreignKey:OriginId;references:ZoneId"`
	FareRulesDestinationID     FareRules           `gorm:"foreignKey:DestinationId;references:ZoneId"`
	FareRulesContainsId        FareRules           `gorm:"foreignKey:ContainsId;references:ZoneId"`
	FareLegJoinRulesFromStopID FareLegJoinRules    `gorm:"foreignKey:FromStopId;references:StopId"`
	FareLegJoinRulesToStopID   FareLegJoinRules    `gorm:"foreignKey:ToStopId;references:StopId"`
	StopArea                   []StopArea          `gorm:"foreignKey:StopId;references:StopId "`
	TransferFromStopID         []Transfer          `gorm:"foreignKey:FromStopId;references:StopId"`
	TransferToStopID           []Transfer          `gorm:"foreignKey:ToStopId;references:StopId"`
	PathwayFromStopID          []Pathway           `gorm:"foreignKey:FromStopId;references:StopId"`
	PathwayToStopID            []Pathway           `gorm:"foreignKey:ToStopId;references:StopId"`
	LocationGroupStop          []LocationGroupStop `gorm:"foreignKey:StopId;references:StopId "`
}

func ParseStop(path string) ([]Stop, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stops CSV: %w", err)
	}

	// データを解析してStop構造体のスライスを作成
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

		ttsStopName, err := df.GetStringPtr(i, "tts_stop_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'tts_stop_name' at row %d: %w", i, err)
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

		zoneId, err := df.GetString(i, "zone_id")
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

		parentStation, err := df.GetIntPtr(i, "parent_station")
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

		// Stop構造体を作成しリストに追加
		stops = append(stops, Stop{
			StopId:             stopId,
			StopCode:           stopCode,
			StopName:           stopName,
			TtsStopName:        ttsStopName,
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
