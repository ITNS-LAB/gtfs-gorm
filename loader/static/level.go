package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func LoadLevels(path string) []static.Level {
	var levels []static.Level
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		levels = append(levels, static.Level{
			LevelId:    dataframe.IsBlank(df.GetElement("level_id")),
			LevelIndex: dataframe.ParseFloat64(df.GetElement("level_index")),
			LevelName:  dataframe.IsBlank(df.GetElement("level_name")),
		})
	}
	return levels
}
