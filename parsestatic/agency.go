package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseAgency(path string) ([]gtfsjp.Agency, error) {
	var agencies []gtfsjp.Agency
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

		agencies = append(agencies, gtfsjp.Agency{
			AgencyId:       util.IsBlank(df.GetElement("agency_id")),
			AgencyName:     util.IsBlank(df.GetElement("agency_name")),
			AgencyUrl:      util.IsBlank(df.GetElement("agency_url")),
			AgencyTimezone: util.IsBlank(df.GetElement("agency_timezone")),
			AgencyLang:     util.IsBlank(df.GetElement("agency_lang")),
			AgencyPhone:    util.IsBlank(df.GetElement("agency_phone")),
			AgencyFareUrl:  util.IsBlank(df.GetElement("agency_fare_url")),
			AgencyEmail:    util.IsBlank(df.GetElement("agency_email")),
		})
	}
	return agencies, nil
}
