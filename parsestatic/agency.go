package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseAgency(path string) []ormstatic.Agency {
	var agencies []ormstatic.Agency
	df := dataframe.OpenCsv(path)
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
	return agencies
}
