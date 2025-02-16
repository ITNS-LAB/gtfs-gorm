package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Areas struct {
	AreaId   string `gorm:"primaryKey"`
	AreaName *string
	//FareLegFromAreaID []FareLeg  `gorm:"foreignKey:FromAreaId;references:AreaId"`
	//FareLegToAreaID   []FareLeg  `gorm:"foreignKey:ToAreaId;references:AreaId"`
	StopArea []StopArea `gorm:"foreignKey:AreaId;references:AreaId"`
}

func ParseAreas(path string) ([]Areas, error) {
	//CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open Areas CSV: %w", err)
	}

	//データを解析してAgency構造体のスライスを作成
	var areass []Areas
	for i := 0; i < len(df.Records); i++ {
		areaID, err := df.GetString(i, "areaID")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'areaID' at row %d: %w", i, err)
		}

		areaName, err := df.GetStringPtr(i, "areaName")
		if err != nil {

			return nil, fmt.Errorf("failed to get 'areaName' at row %d: %w", i, err)
		}

		//Agency 構造体を作成しリストに追加
		areass = append(areass, Areas{
			AreaId:   areaID,
			AreaName: areaName,
		})
	}

	return areass, nil
}
