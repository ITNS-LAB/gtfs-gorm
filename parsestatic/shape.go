package parsestatic

import (
	"fmt"
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
			fmt.Println("Error:", err)
			break
		}

		shapes = append(shapes, ormstatic.Shape{
			ShapeId:           dataframe.IsBlank(df.GetElement("shape_id")),
			ShapePtLat:        dataframe.ParseFloat64(df.GetElement("shape_pt_lat")),
			ShapePtLon:        dataframe.ParseFloat64(df.GetElement("shape_pt_lon")),
			ShapePtSequence:   dataframe.ParseInt(df.GetElement("shape_pt_sequence")),
			ShapeDistTraveled: dataframe.ParseFloat64(df.GetElement("shape_dist_traveled")),
		})
	}
	return shapes, nil
}
