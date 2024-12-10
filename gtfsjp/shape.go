package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"github.com/paulmach/orb"
)

type Shape struct {
	ShapeId           string  `gorm:"primaryKey"`
	ShapePtLat        float64 `gorm:"not null"`
	ShapePtLon        float64 `gorm:"not null"`
	ShapePtSequence   int     `gorm:"primaryKey"`
	ShapeDistTraveled *float64
}

func (Shape) TableName() string {
	return "shapes"
}

func ParseShapes(path string) ([]Shape, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open shapes CSV: %w", err)
	}

	// データを解析して Shape 構造体のスライスを作成
	var shapes []Shape
	for i := 0; i < len(df.Records); i++ {
		shapeId, err := df.GetString(i, "shape_id")
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
			return nil, fmt.Errorf("failed to get 'shape_dist_traveled' at row %d: %w", i, err)
		}

		// Shape 構造体を作成しリストに追加
		shapes = append(shapes, Shape{
			ShapeId:           shapeId,
			ShapePtLat:        shapePtLat,
			ShapePtLon:        shapePtLon,
			ShapePtSequence:   shapePtSequence,
			ShapeDistTraveled: shapeDistTraveled,
		})
	}

	return shapes, nil
}

type ShapeGeom struct {
	ShapeId           string  `gorm:"primaryKey"`
	ShapePtLat        float64 `gorm:"not null"`
	ShapePtLon        float64 `gorm:"not null"`
	ShapePtSequence   int     `gorm:"primaryKey"`
	ShapeDistTraveled *float64
	Geom              gormdatatypes.Geometry `gorm:"index"`
}

func (ShapeGeom) TableName() string {
	return "shapes_geom"
}

func ParseShapesGeom(path string) ([]ShapeGeom, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open shapes CSV: %w", err)
	}

	// データを解析して Shape 構造体のスライスを作成
	var shapesGeom []ShapeGeom
	for i := 0; i < len(df.Records); i++ {
		shapeId, err := df.GetString(i, "shape_id")
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
			return nil, fmt.Errorf("failed to get 'shape_dist_traveled' at row %d: %w", i, err)
		}

		geom := gormdatatypes.Geometry{
			Geom: orb.Point{shapePtLon, shapePtLat},
			Srid: 4326,
		}

		// Shape 構造体を作成しリストに追加
		shapesGeom = append(shapesGeom, ShapeGeom{
			ShapeId:           shapeId,
			ShapePtLat:        shapePtLat,
			ShapePtLon:        shapePtLon,
			ShapePtSequence:   shapePtSequence,
			ShapeDistTraveled: shapeDistTraveled,
			Geom:              geom,
		})
	}

	return shapesGeom, nil
}
