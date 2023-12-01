package ormstatic

import "gorm.io/datatypes"

type Calendar struct {
	ServiceId *string         `gorm:"primaryKey;index;not null"`
	Monday    *int            `gorm:"not null"`
	Tuesday   *int            `gorm:"not null"`
	Wednesday *int            `gorm:"not null"`
	Thursday  *int            `gorm:"not null"`
	Friday    *int            `gorm:"not null"`
	Saturday  *int            `gorm:"not null"`
	Sunday    *int            `gorm:"not null"`
	StartDate *datatypes.Date `gorm:"not null"`
	EndDate   *datatypes.Date `gorm:"not null"`
	Trip      Trip            `gorm:"foreignKey:ServiceId"`
}

func (Calendar) TableName() string {
	return "calendar"
}
