package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
)

func ParseAgency(path string) ([]ormstatic.Agency, error) {
	var agencies []ormstatic.Agency
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return agencies, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		agencies = append(agencies, ormstatic.Agency{
			AgencyId:       dataframe.IsBlank(df.GetElement("agency_id")),
			AgencyName:     dataframe.IsBlank(df.GetElement("agency_name")),
			AgencyUrl:      dataframe.IsBlank(df.GetElement("agency_url")),
			AgencyTimezone: dataframe.IsBlank(df.GetElement("agency_timezone")),
			AgencyLang:     dataframe.IsBlank(df.GetElement("agency_lang")),
			AgencyPhone:    dataframe.IsBlank(df.GetElement("agency_phone")),
			AgencyFareUrl:  dataframe.IsBlank(df.GetElement("agency_fare_url")),
			AgencyEmail:    dataframe.IsBlank(df.GetElement("agency_email")),
		})
	}
	return agencies, nil
}
