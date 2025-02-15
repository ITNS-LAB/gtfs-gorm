package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Levels struct {
	LevelId    string  `gorm:"primaryKey"`
	LevelIndex float64 `gorm:"not null"`
	LevelName  *string
	Stop       []Stop `gorm:"foreignKey:LevelId;references:LevelId "`
}

func ParseLevels(path string) ([]Levels, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open levels CSV: %w", err)
	}

	// Parse the CSV data into a slice of Levels structs
	var levels []Levels
	for i := 0; i < len(df.Records); i++ {
		levelID, err := df.GetString(i, "level_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'level_id' at row %d: %w", i, err)
		}

		levelIndex, err := df.GetFloat(i, "level_index")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'level_index' at row %d: %w", i, err)
		}

		levelName, err := df.GetStringPtr(i, "level_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'level_name' at row %d: %w", i, err)
		}

		// Create Levels struct and append to the list
		levels = append(levels, Levels{
			LevelId:    levelID,
			LevelIndex: levelIndex,
			LevelName:  levelName,
		})
	}

	return levels, nil
}
