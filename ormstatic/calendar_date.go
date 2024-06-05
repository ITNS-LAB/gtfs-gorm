package ormstatic

import (
	"gorm.io/datatypes"
)

type CalendarDate struct {
	ServiceId     string         `gorm:"primaryKey;index;not null"`
	Date          datatypes.Date `gorm:"primaryKey;index;not null"`
	ExceptionType int16          `gorm:"not null"`
	//UniversalCalendar UniversalCalendar `gorm:"foreignKey:ServiceId"`
}

func (CalendarDate) TableName() string {
	return "calendar_dates"
}

func (c CalendarDate) GetServiceId() any {
	return c.ServiceId
}

func (c CalendarDate) GetDate() any {
	return c.Date
}

func (c CalendarDate) GetExceptionType() any {
	return c.ExceptionType
}
