package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/paulmach/orb"
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

		point := orb.Point{*util.ParseFloat64(df.GetElement("shape_pt_lon")),
			*util.ParseFloat64(df.GetElement("shape_pt_lat"))}

		shapes = append(shapes, ormstatic.Shape{
			ShapeId:           util.IsBlank(df.GetElement("shape_id")),
			ShapePtLat:        util.ParseFloat64(df.GetElement("shape_pt_lat")),
			ShapePtLon:        util.ParseFloat64(df.GetElement("shape_pt_lon")),
			ShapePtSequence:   util.ParseInt(df.GetElement("shape_pt_sequence")),
			ShapeDistTraveled: util.ParseFloat64(df.GetElement("shape_dist_traveled")),
			Geom:              &geomdatatypes.Geometry{Geom: point, Srid: 4326},
		})
	}
	return shapes, nil
}
