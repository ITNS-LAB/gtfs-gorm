package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type PatternJp struct {
	JpPatternId     string `gorm:"primaryKey"`
	RouteUpdateDate *datatypes.Date
	OriginStop      *string
	ViaStop         *string
	DestinationStop *string
}

func (PatternJp) TableName() string {
	return "pattern_jp"
}

func ParsePatternJp(path string) ([]PatternJp, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open pattern_jp CSV: %w", err)
	}

	// データを解析して PatternJp 構造体のスライスを作成
	var patterns []PatternJp
	for i := 0; i < len(df.Records); i++ {
		jpPatternId, err := df.GetString(i, "jp_pattern_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_pattern_id' at row %d: %w", i, err)
		}

		routeUpdateDate, err := df.GetDatePtr(i, "route_update_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_update_date' at row %d: %w", i, err)
		}

		originStop, err := df.GetStringPtr(i, "origin_stop")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'origin_stop' at row %d: %w", i, err)
		}

		viaStop, err := df.GetStringPtr(i, "via_stop")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'via_stop' at row %d: %w", i, err)
		}

		destinationStop, err := df.GetStringPtr(i, "destination_stop")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'destination_stop' at row %d: %w", i, err)
		}

		// PatternJp 構造体を作成しリストに追加
		patterns = append(patterns, PatternJp{
			JpPatternId:     jpPatternId,
			RouteUpdateDate: routeUpdateDate,
			OriginStop:      originStop,
			ViaStop:         viaStop,
			DestinationStop: destinationStop,
		})
	}

	return patterns, nil
}

type PatternJpGeom struct {
	JpPatternId     string `gorm:"primaryKey"`
	RouteUpdateDate *datatypes.Date
	OriginStop      *string
	ViaStop         *string
	DestinationStop *string
}

func (PatternJpGeom) TableName() string {
	return "pattern_jp"
}

func ParsePatternJpGeom(path string) ([]PatternJpGeom, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open pattern_jp CSV: %w", err)
	}

	// データを解析して PatternJp 構造体のスライスを作成
	var patterns []PatternJpGeom
	for i := 0; i < len(df.Records); i++ {
		jpPatternId, err := df.GetString(i, "jp_pattern_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'jp_pattern_id' at row %d: %w", i, err)
		}

		routeUpdateDate, err := df.GetDatePtr(i, "route_update_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_update_date' at row %d: %w", i, err)
		}

		originStop, err := df.GetStringPtr(i, "origin_stop")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'origin_stop' at row %d: %w", i, err)
		}

		viaStop, err := df.GetStringPtr(i, "via_stop")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'via_stop' at row %d: %w", i, err)
		}

		destinationStop, err := df.GetStringPtr(i, "destination_stop")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'destination_stop' at row %d: %w", i, err)
		}

		// PatternJp 構造体を作成しリストに追加
		patterns = append(patterns, PatternJpGeom{
			JpPatternId:     jpPatternId,
			RouteUpdateDate: routeUpdateDate,
			OriginStop:      originStop,
			ViaStop:         viaStop,
			DestinationStop: destinationStop,
		})
	}

	return patterns, nil
}
