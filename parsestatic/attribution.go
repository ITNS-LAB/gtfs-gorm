package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
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
			return attributions, err
		}

		organizationName, err := dataframe.ParseString(df.GetElement("organization_name"))
		if err != nil {
			return []ormstatic.Attribution{}, err
		}
		isProducer, err := dataframe.ParseNullInt16(df.GetElement("is_producer"))
		isOperator, err := dataframe.ParseNullInt16(df.GetElement("is_operator"))
		isAuthority, err := dataframe.ParseNullInt16(df.GetElement("is-authority"))

		attributions = append(attributions, ormstatic.Attribution{
			AttributionId:    df.GetElement("attribution_id"),
			AgencyId:         df.GetElement("agency_id"),
			RouteId:          df.GetElement("route_id"),
			TripId:           df.GetElement("trip_id"),
			OrganizationName: organizationName,
			IsProducer:       isProducer,
			IsOperator:       isOperator,
			IsAuthority:      isAuthority,
			AttributionUrl:   df.GetElement("attribution_url"),
			AttributionEmail: df.GetElement("attribution_email"),
			AttributionPhone: df.GetElement("attribution_phone"),
		})
	}
	return attributions, nil
}
