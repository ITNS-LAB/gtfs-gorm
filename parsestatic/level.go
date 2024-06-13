package parsestatic

import (
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
			return []ormstatic.Level{}, err
		}

		levelId, err := dataframe.ParseString(df.GetElement("level_id"))
		if err != nil {
			return []ormstatic.Level{}, err
		}

		levelIndex, err := dataframe.ParseFloat64(df.GetElement("level_index"))
		if err != nil {
			return []ormstatic.Level{}, err
		}

		levels = append(levels, ormstatic.Level{
			LevelId:    levelId,
			LevelIndex: levelIndex,
			LevelName:  df.GetElement("level_name"),
		})
	}
	return levels, nil
}
