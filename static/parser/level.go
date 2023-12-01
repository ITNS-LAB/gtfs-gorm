package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseLevels(path string) []orm.Level {
	var levels []orm.Level
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		levels = append(levels, orm.Level{
			LevelId:    dataframe.IsBlank(df.GetElement("level_id")),
			LevelIndex: dataframe.ParseFloat64(df.GetElement("level_index")),
			LevelName:  dataframe.IsBlank(df.GetElement("level_name")),
		})
	}
	return levels
}
