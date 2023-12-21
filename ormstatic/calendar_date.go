package ormstatic

import "gorm.io/datatypes"

type CalendarDate struct {
	ServiceId         *string           `gorm:"primaryKey;index;not null"`
	Date              *datatypes.Date   `gorm:"primaryKey;index;not null"`
	ExceptionType     *int              `gorm:"not null"`
	UniversalCalendar UniversalCalendar `gorm:"foreignKey:ServiceId"`
}

func (CalendarDate) TableName() string {
	return "calendar_dates"
}
