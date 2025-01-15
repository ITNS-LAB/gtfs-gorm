package gtfsschedule

type Levels struct {
	LevelID    string  `gorm:"primaryKey"`
	LevelIndex float64 `gorm:"not null"`
	LevelName  *string
}
