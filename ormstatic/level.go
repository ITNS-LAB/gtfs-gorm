package ormstatic

import "database/sql"

type Level struct {
	LevelId    string  `gorm:"primaryKey;index;not null"`
	LevelIndex float64 `gorm:"index;not null"`
	LevelName  sql.NullString
}

func (Level) TableName() string {
	return "levels"
}

func (l Level) GetLevelId() any {
	return l.LevelId
}

func (l Level) GetLevelIndex() any {
	return l.LevelIndex
}

func (l Level) GetLevelName() any {
	if l.LevelName.Valid {
		return l.LevelName.String
	}
	return nil
}
