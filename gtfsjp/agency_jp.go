package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/csvutil"
)

type AgencyJp struct {
	AgencyId            string `gorm:"primaryKey"`
	AgencyOfficialName  *string
	AgencyZipNumber     *string
	AgencyAddress       *string
	AgencyPresidentPos  *string
	AgencyPresidentName *string
}

func (AgencyJp) TableName() string {
	return "agency_jp"
}

func ParseAgencyJp(path string) ([]AgencyJp, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open agency_jp CSV: %w", err)
	}

	// データを解析して AgencyJp 構造体のスライスを作成
	var agenciesJp []AgencyJp
	for i := 0; i < len(df.Records); i++ { // df.records を使用
		agencyId, err := df.GetString(i, "agency_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_id' at row %d: %w", i, err)
		}

		agencyOfficialName, err := df.GetStringPtr(i, "agency_official_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_official_name' at row %d: %w", i, err)
		}

		agencyZipNumber, err := df.GetStringPtr(i, "agency_zip_number")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_zip_number' at row %d: %w", i, err)
		}

		agencyAddress, err := df.GetStringPtr(i, "agency_address")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_address' at row %d: %w", i, err)
		}

		agencyPresidentPos, err := df.GetStringPtr(i, "agency_president_pos")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_president_pos' at row %d: %w", i, err)
		}

		agencyPresidentName, err := df.GetStringPtr(i, "agency_president_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'agency_president_name' at row %d: %w", i, err)
		}

		// AgencyJp 構造体を作成しリストに追加
		agenciesJp = append(agenciesJp, AgencyJp{
			AgencyId:            agencyId,
			AgencyOfficialName:  agencyOfficialName,
			AgencyZipNumber:     agencyZipNumber,
			AgencyAddress:       agencyAddress,
			AgencyPresidentPos:  agencyPresidentPos,
			AgencyPresidentName: agencyPresidentName,
		})
	}

	return agenciesJp, nil
}
