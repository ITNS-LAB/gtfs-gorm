package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseLevels(path string) ([]gtfsjp.Level, error) {
	var levels []gtfsjp.Level
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

		levels = append(levels, gtfsjp.Level{
			LevelId:    util.IsBlank(df.GetElement("level_id")),
			LevelIndex: util.ParseFloat64(df.GetElement("level_index")),
			LevelName:  util.IsBlank(df.GetElement("level_name")),
		})
	}
	return levels, nil
}
