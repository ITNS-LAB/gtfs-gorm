package ormstatic

import "gorm.io/datatypes"

type UniversalCalendar struct {
	ServiceId *string         `gorm:"primaryKey;index"`
	Date      *datatypes.Date `gorm:"primaryKey;index"`
}

func (UniversalCalendar) TableName() string {
	return "universal_calendar"
}
