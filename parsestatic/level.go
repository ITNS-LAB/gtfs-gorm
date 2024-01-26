package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseLevels(path string) ([]ormstatic.Level, error) {
	var levels []ormstatic.Level
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return levels, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		levels = append(levels, ormstatic.Level{
			LevelId:    dataframe.IsBlank(df.GetElement("level_id")),
			LevelIndex: dataframe.ParseFloat64(df.GetElement("level_index")),
			LevelName:  dataframe.IsBlank(df.GetElement("level_name")),
		})
	}
	return levels, nil
}
