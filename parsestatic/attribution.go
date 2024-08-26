package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
)

func ParseAttributions(path string) ([]ormstatic.Attribution, error) {
	var attributions []ormstatic.Attribution
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

		attributions = append(attributions, ormstatic.Attribution{
			AttributionId:    dataframe.IsBlank(df.GetElement("attribution_id")),
			AgencyId:         dataframe.IsBlank(df.GetElement("agency_id")),
			RouteId:          dataframe.IsBlank(df.GetElement("route_id")),
			TripId:           dataframe.IsBlank(df.GetElement("trip_id")),
			OrganizationName: dataframe.IsBlank(df.GetElement("organization_name")),
			IsProducer:       dataframe.ParseEnum(df.GetElement("is_producer")),
			IsOperator:       dataframe.ParseEnum(df.GetElement("is_operator")),
			IsAuthority:      dataframe.ParseEnum(df.GetElement("is_authority")),
			AttributionUrl:   dataframe.IsBlank(df.GetElement("attribution_url")),
			AttributionEmail: dataframe.IsBlank(df.GetElement("attribution_email")),
			AttributionPhone: dataframe.IsBlank(df.GetElement("attribution_phone")),
		})
	}
	return attributions, nil
}
