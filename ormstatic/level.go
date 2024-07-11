package ormstatic

type Level struct {
	LevelId    *string  `gorm:"primaryKey"`
	LevelIndex *float64 `gorm:"index;not null"`
	LevelName  *string
}

func (Level) TableName() string {
	return "levels"
}
