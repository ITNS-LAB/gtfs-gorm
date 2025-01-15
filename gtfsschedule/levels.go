package gtfsschedule

type Levels struct {
	LevelID    string  `gorm:"primaryKey"`
	LevelIndex float64 `gorm:"not null"`
	LevelName  *string
	Stop       []Stop `gorm:"foreignKey:LevelID;references:LevelId "`
}
