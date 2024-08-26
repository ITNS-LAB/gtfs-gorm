package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
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
			fmt.Println("Error:", err)
			break
		}

		fareRules = append(fareRules, ormstatic.FareRule{
			FareId:        dataframe.IsBlank(df.GetElement("fare_id")),
			RouteId:       dataframe.IsBlank(df.GetElement("route_id")),
			OriginId:      dataframe.IsBlank(df.GetElement("origin_id")),
			DestinationId: dataframe.IsBlank(df.GetElement("destination_id")),
			ContainsId:    dataframe.IsBlank(df.GetElement("contains_id")),
		})
	}
	return fareRules, nil
}
