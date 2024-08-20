package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	geomdatatypes "github.com/ITNS-LAB/gtfs-gorm/pkg/gormdatatypes"
	"github.com/paulmach/orb"
)

func ParseStops(path string) ([]ormstatic.Stop, error) {
	var stops []ormstatic.Stop
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return stops, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		point := orb.Point{*dataframe.ParseFloat64(df.GetElement("stop_lon")),
			*dataframe.ParseFloat64(df.GetElement("stop_lat"))}

		stops = append(stops, ormstatic.Stop{
			StopId:             dataframe.IsBlank(df.GetElement("stop_id")),
			StopCode:           dataframe.IsBlank(df.GetElement("stop_code")),
			StopName:           dataframe.IsBlank(df.GetElement("stop_name")),
			StopDesc:           dataframe.IsBlank(df.GetElement("stop_desc")),
			StopLat:            dataframe.ParseFloat64(df.GetElement("stop_lat")),
			StopLon:            dataframe.ParseFloat64(df.GetElement("stop_lon")),
			ZoneId:             dataframe.IsBlank(df.GetElement("zone_id")),
			StopUrl:            dataframe.IsBlank(df.GetElement("stop_url")),
			LocationType:       dataframe.ParseEnum(df.GetElement("location_type")),
			ParentStation:      dataframe.IsBlank(df.GetElement("parent_station")),
			StopTimezone:       dataframe.IsBlank(df.GetElement("stop_timezone")),
			WheelchairBoarding: dataframe.ParseEnum(df.GetElement("wheelchair_boarding")),
			LevelId:            dataframe.IsBlank(df.GetElement("level_id")),
			PlatformCode:       dataframe.IsBlank(df.GetElement("platform_code")),
			Geom:               &geomdatatypes.Geometry{Geom: point, Srid: 4326},
		})
	}
	return stops, nil
}
