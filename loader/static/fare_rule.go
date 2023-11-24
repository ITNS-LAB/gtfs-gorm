package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func LoadFareRules(path string) []static.FareRule {
	var fareRules []static.FareRule
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fareRules = append(fareRules, static.FareRule{
			FareId:        dataframe.IsBlank(df.GetElement("fare_id")),
			RouteId:       dataframe.IsBlank(df.GetElement("route_id")),
			OriginId:      dataframe.IsBlank(df.GetElement("origin_id")),
			DestinationId: dataframe.IsBlank(df.GetElement("destination_id")),
			ContainsId:    dataframe.IsBlank(df.GetElement("contains_id")),
		})
	}
	return fareRules
}
