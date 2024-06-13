package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
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
			return []ormstatic.Agency{}, err
		}

		agencyId, err := dataframe.ParseString(df.GetElement("agency_id"))
		if err != nil {
			return []ormstatic.Agency{}, err
		}

		agencyName, err := dataframe.ParseString(df.GetElement("agency_name"))
		if err != nil {
			return []ormstatic.Agency{}, err
		}

		agencyUrl, err := dataframe.ParseString(df.GetElement("agency_url"))
		if err != nil {
			return []ormstatic.Agency{}, err
		}

		agencyTimeZone, err := dataframe.ParseString(df.GetElement("agency_timezone"))
		if err != nil {
			return []ormstatic.Agency{}, err
		}

		agencies = append(agencies, ormstatic.Agency{
			AgencyId:       agencyId,
			AgencyName:     agencyName,
			AgencyUrl:      agencyUrl,
			AgencyTimezone: agencyTimeZone,
			AgencyLang:     df.GetElement("agency_lang"),
			AgencyPhone:    df.GetElement("agency_phone"),
			AgencyFareUrl:  df.GetElement("agency_fare_url"),
			AgencyEmail:    df.GetElement("agency_email"),
		})
	}
	return agencies, nil
}
