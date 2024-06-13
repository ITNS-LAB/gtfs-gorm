package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseFareRules(path string) ([]ormstatic.FareRule, error) {
	var fareRules []ormstatic.FareRule
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return fareRules, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			return []ormstatic.FareRule{}, err
		}

		fareId, err := dataframe.ParseString(df.GetElement("fare_id"))
		if err != nil {
			return []ormstatic.FareRule{}, err
		}

		fareRules = append(fareRules, ormstatic.FareRule{
			FareId:        fareId,
			RouteId:       df.GetElement("route_id"),
			OriginId:      df.GetElement("origin_id"),
			DestinationId: df.GetElement("destination_id"),
			ContainsId:    df.GetElement("contains_id"),
		})
	}
	return fareRules, nil
}
