package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/csvutil"
)

type OfficeJp struct {
	OfficeId    string `gorm:"primaryKey"`
	OfficeName  string `gorm:"not null"`
	OfficeUrl   *string
	OfficePhone *string
}

func (OfficeJp) TableName() string {
	return "office_jp"
}

func ParseOfficeJp(path string) ([]OfficeJp, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open office_jp CSV: %w", err)
	}

	// データを解析して OfficeJp 構造体のスライスを作成
	var offices []OfficeJp
	for i := 0; i < len(df.Records); i++ {
		officeId, err := df.GetString(i, "office_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'office_id' at row %d: %w", i, err)
		}

		officeName, err := df.GetString(i, "office_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'office_name' at row %d: %w", i, err)
		}

		officeUrl, err := df.GetStringPtr(i, "office_url")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'office_url' at row %d: %w", i, err)
		}

		officePhone, err := df.GetStringPtr(i, "office_phone")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'office_phone' at row %d: %w", i, err)
		}

		// OfficeJp 構造体を作成しリストに追加
		offices = append(offices, OfficeJp{
			OfficeId:    officeId,
			OfficeName:  officeName,
			OfficeUrl:   officeUrl,
			OfficePhone: officePhone,
		})
	}

	return offices, nil
}
