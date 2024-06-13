package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
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
			return []ormstatic.Stop{}, err
		}

		stopId, err := dataframe.ParseString(df.GetElement("stop_id"))
		if err != nil {
			return []ormstatic.Stop{}, err
		}

		stopLat, err := dataframe.ParseNullFloat64(df.GetElement("stop_lat"))
		stopLon, err := dataframe.ParseNullFloat64(df.GetElement("stop_lon"))
		locationType, err := dataframe.ParseNullInt16(df.GetElement("location_type"))
		wheelchairBoarding, err := dataframe.ParseNullInt16(df.GetElement("wheelchair_boarding"))

		stops = append(stops, ormstatic.Stop{
			StopId:             stopId,
			StopCode:           df.GetElement("stop_code"),
			StopName:           df.GetElement("stop_name"),
			StopDesc:           df.GetElement("stop_desc"),
			StopLat:            stopLat,
			StopLon:            stopLon,
			ZoneId:             df.GetElement("zone_id"),
			StopUrl:            df.GetElement("stop_url"),
			LocationType:       locationType,
			ParentStation:      df.GetElement("parent_station"),
			StopTimezone:       df.GetElement("stop_timezone"),
			WheelchairBoarding: wheelchairBoarding,
			LevelId:            df.GetElement("level_id"),
			PlatformCode:       df.GetElement("platform_code"),
		})
	}
	return stops, nil
}
