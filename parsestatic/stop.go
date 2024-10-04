package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
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

		point := orb.Point{*util.ParseFloat64(df.GetElement("stop_lon")),
			*util.ParseFloat64(df.GetElement("stop_lat"))}

		stops = append(stops, ormstatic.Stop{
			StopId:             util.IsBlank(df.GetElement("stop_id")),
			StopCode:           util.IsBlank(df.GetElement("stop_code")),
			StopName:           util.IsBlank(df.GetElement("stop_name")),
			StopDesc:           util.IsBlank(df.GetElement("stop_desc")),
			StopLat:            util.ParseFloat64(df.GetElement("stop_lat")),
			StopLon:            util.ParseFloat64(df.GetElement("stop_lon")),
			ZoneId:             util.IsBlank(df.GetElement("zone_id")),
			StopUrl:            util.IsBlank(df.GetElement("stop_url")),
			LocationType:       util.ParseEnum(df.GetElement("location_type")),
			ParentStation:      util.IsBlank(df.GetElement("parent_station")),
			StopTimezone:       util.IsBlank(df.GetElement("stop_timezone")),
			WheelchairBoarding: util.ParseEnum(df.GetElement("wheelchair_boarding")),
			LevelId:            util.IsBlank(df.GetElement("level_id")),
			PlatformCode:       util.IsBlank(df.GetElement("platform_code")),
			Geom:               &geomdatatypes.Geometry{Geom: point, Srid: 4326},
		})
	}
	return stops, nil
}
