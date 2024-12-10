package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/csvutil"
)

type Agency struct {
	//Id             int    `gorm:"primaryKey;auto_increment"`
	AgencyId       string `gorm:"primaryKey"`
	AgencyName     string `gorm:"not null"`
	AgencyUrl      string `gorm:"not null"`
	AgencyTimezone string `gorm:"not null;default:Asia/Tokyo"`
	AgencyLang     string `gorm:"not null;default:ja"`
	AgencyPhone    *string
	AgencyFareUrl  *string
	AgencyEmail    *string
	AgencyJp       AgencyJp `gorm:"foreignKey:AgencyId"`
	Routes         []Route  `gorm:"foreignKey:AgencyId;references:AgencyId"`
}

func (Agency) TableName() string {
	return "agency"
}

func ParseAgency(path string) ([]Agency, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open agency CSV: %w", err)
	}

	// データを解析して Agency 構造体のスライスを作成
	var agencies []Agency
	for i := 0; i < len(df.Records); i++ {
		agencyId, err := df.GetString(i, "agency_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_id' at row %d: %w", i, err)
		}

		agencyName, err := df.GetString(i, "agency_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_name' at row %d: %w", i, err)
		}

		agencyUrl, err := df.GetString(i, "agency_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_url' at row %d: %w", i, err)
		}

		agencyTimezone, err := df.GetString(i, "agency_timezone")
		if err != nil {
			agencyTimezone = "Asia/Tokyo" // デフォルト値
		}

		agencyLang, err := df.GetString(i, "agency_lang")
		if err != nil {
			agencyLang = "ja" // デフォルト値
		}

		agencyPhone, err := df.GetStringPtr(i, "agency_phone")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_phone' at row %d: %w", i, err)
		}

		agencyFareUrl, err := df.GetStringPtr(i, "agency_fare_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_fare_url' at row %d: %w", i, err)
		}

		agencyEmail, err := df.GetStringPtr(i, "agency_email")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_email' at row %d: %w", i, err)
		}

		// Agency 構造体を作成しリストに追加
		agencies = append(agencies, Agency{
			AgencyId:       agencyId,
			AgencyName:     agencyName,
			AgencyUrl:      agencyUrl,
			AgencyTimezone: agencyTimezone,
			AgencyLang:     agencyLang,
			AgencyPhone:    agencyPhone,
			AgencyFareUrl:  agencyFareUrl,
			AgencyEmail:    agencyEmail,
		})
	}

	return agencies, nil
}
