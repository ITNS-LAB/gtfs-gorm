package gtfsschedule

import "gorm.io/datatypes"

type CalendarDates struct {
	ServiceID     int            `gorm:"primary_key"`
	Date          datatypes.Date `gorm:"not null"`
	ExceptionType int            `gorm:"not null"`
}
