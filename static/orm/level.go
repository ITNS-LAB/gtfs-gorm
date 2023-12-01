package orm

type Level struct {
	LevelId    *string  `gorm:"primaryKey;index;not null"`
	LevelIndex *float64 `gorm:"index;not null"`
	LevelName  *string
	Stop       Stop `gorm:"foreignKey:LevelId"`
}

func (Level) TableName() string {
	return "levels"
}
