package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseShapes(path string) []static.Shape {
	var shapes []static.Shape
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		shapes = append(shapes, static.Shape{
			ShapeId:           dataframe.IsBlank(df.GetElement("shape_id")),
			ShapePtLat:        dataframe.ParseFloat64(df.GetElement("shape_pt_lat")),
			ShapePtLon:        dataframe.ParseFloat64(df.GetElement("shape_pt_lon")),
			ShapePtSequence:   dataframe.ParseInt(df.GetElement("shape_pt_sequence")),
			ShapeDistTraveled: dataframe.ParseFloat64(df.GetElement("shape_dist_traveled")),
		})
	}
	return shapes
}
