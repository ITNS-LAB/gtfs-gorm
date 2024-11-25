package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseAttributions(path string) ([]gtfsjp.Attribution, error) {
	var attributions []gtfsjp.Attribution
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return attributions, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		attributions = append(attributions, gtfsjp.Attribution{
			AttributionId:    util.IsBlank(df.GetElement("attribution_id")),
			AgencyId:         util.IsBlank(df.GetElement("agency_id")),
			RouteId:          util.IsBlank(df.GetElement("route_id")),
			TripId:           util.IsBlank(df.GetElement("trip_id")),
			OrganizationName: util.IsBlank(df.GetElement("organization_name")),
			IsProducer:       util.ParseEnum(df.GetElement("is_producer")),
			IsOperator:       util.ParseEnum(df.GetElement("is_operator")),
			IsAuthority:      util.ParseEnum(df.GetElement("is_authority")),
			AttributionUrl:   util.IsBlank(df.GetElement("attribution_url")),
			AttributionEmail: util.IsBlank(df.GetElement("attribution_email")),
			AttributionPhone: util.IsBlank(df.GetElement("attribution_phone")),
		})
	}
	return attributions, nil
}
