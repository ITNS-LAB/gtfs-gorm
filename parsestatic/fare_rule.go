package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseFareRules(path string) ([]gtfsjp.FareRule, error) {
	var fareRules []gtfsjp.FareRule
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

		fareRules = append(fareRules, gtfsjp.FareRule{
			FareId:        util.IsBlank(df.GetElement("fare_id")),
			RouteId:       util.IsBlank(df.GetElement("route_id")),
			OriginId:      util.IsBlank(df.GetElement("origin_id")),
			DestinationId: util.IsBlank(df.GetElement("destination_id")),
			ContainsId:    util.IsBlank(df.GetElement("contains_id")),
		})
	}
	return fareRules, nil
}
