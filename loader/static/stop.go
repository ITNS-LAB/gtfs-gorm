package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func LoadStops(path string) []static.Stop {
	var stops []static.Stop
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		stops = append(stops, static.Stop{
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
		})
	}
	return stops
}
