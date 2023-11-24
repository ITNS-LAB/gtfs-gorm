package static

import "gorm.io/datatypes"

type CalendarDate struct {
	ServiceId     *string         `gorm:"primaryKey;index;not null"`
	Date          *datatypes.Date `gorm:"primaryKey;index;not null"`
	ExceptionType *int            `gorm:"not null"`
	Calendar      Calendar        `gorm:"foreignKey:ServiceId"`
}

func (CalendarDate) TableName() string {
	return "calendar_dates"
}
