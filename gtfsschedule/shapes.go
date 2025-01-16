package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Shape struct {
	ShapeID           string  `gorm:"primary_key"`
	ShapePtLat        float64 `gorm:"not null"`
	ShapePtLon        float64 `gorm:"not null"`
	ShapePtSequence   int     `gorm:"not null"`
	ShapeDistTraveled *float64
	Trips             []Trips `gorm:"foreignKey:ShapeID;references:ShapeID"`
}

func ParseShapes(path string) ([]Shape, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open Shapes CSV: %w", err)
	}

	// データを解析してShape構造体のスライスを作成
	var shapes []Shape
	for i := 0; i < len(df.Records); i++ {
		shapeID, err := df.GetString(i, "shape_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'shape_id' at row %d: %w", i, err)
		}

		shapePtLat, err := df.GetFloat(i, "shape_pt_lat")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'shape_pt_lat' at row %d: %w", i, err)
		}

		shapePtLon, err := df.GetFloat(i, "shape_pt_lon")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'shape_pt_lon' at row %d: %w", i, err)
		}

		shapePtSequence, err := df.GetInt(i, "shape_pt_sequence")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'shape_pt_sequence' at row %d: %w", i, err)
		}

		shapeDistTraveled, err := df.GetFloatPtr(i, "shape_dist_traveled")
		if err != nil {
			shapeDistTraveled = nil // データがない場合、nilを設定
		}

		// Shape構造体を作成しリストに追加
		shapes = append(shapes, Shape{
			ShapeID:           shapeID,
			ShapePtLat:        shapePtLat,
			ShapePtLon:        shapePtLon,
			ShapePtSequence:   shapePtSequence,
			ShapeDistTraveled: shapeDistTraveled,
		})
	}

	return shapes, nil
}
