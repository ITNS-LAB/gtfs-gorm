package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseShapes(path string) ([]ormstatic.Shape, error) {
	var shapes []ormstatic.Shape
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return shapes, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			return []ormstatic.Shape{}, err
		}

		shapeId, err := dataframe.ParseString(df.GetElement("shape_id"))
		if err != nil {
			return []ormstatic.Shape{}, err
		}

		shapePtLat, err := dataframe.ParseFloat64(df.GetElement("shape_pt_lat"))
		if err != nil {
			return []ormstatic.Shape{}, err
		}

		shapePtLon, err := dataframe.ParseFloat64(df.GetElement("shape_pt_lon"))
		if err != nil {
			return []ormstatic.Shape{}, err
		}

		shapePtSequence, err := dataframe.ParseInt32(df.GetElement("shape_pt_sequence"))
		if err != nil {
			return []ormstatic.Shape{}, err
		}

		shapeDistTraveled, err := dataframe.ParseNullFloat64(df.GetElement("shape_dist_traveled"))

		shapes = append(shapes, ormstatic.Shape{
			ShapeId:           shapeId,
			ShapePtLat:        shapePtLat,
			ShapePtLon:        shapePtLon,
			ShapePtSequence:   shapePtSequence,
			ShapeDistTraveled: shapeDistTraveled,
		})
	}
	return shapes, nil
}
