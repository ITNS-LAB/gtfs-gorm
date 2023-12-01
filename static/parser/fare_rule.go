package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseFareRules(path string) []orm.FareRule {
	var fareRules []orm.FareRule
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fareRules = append(fareRules, orm.FareRule{
			FareId:        dataframe.IsBlank(df.GetElement("fare_id")),
			RouteId:       dataframe.IsBlank(df.GetElement("route_id")),
			OriginId:      dataframe.IsBlank(df.GetElement("origin_id")),
			DestinationId: dataframe.IsBlank(df.GetElement("destination_id")),
			ContainsId:    dataframe.IsBlank(df.GetElement("contains_id")),
		})
	}
	return fareRules
}
