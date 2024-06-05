package ormstatic

import (
	"gorm.io/datatypes"
)

type Calendar struct {
	ServiceId string         `gorm:"primaryKey;index;not null"`
	Monday    int16          `gorm:"not null"`
	Tuesday   int16          `gorm:"not null"`
	Wednesday int16          `gorm:"not null"`
	Thursday  int16          `gorm:"not null"`
	Friday    int16          `gorm:"not null"`
	Saturday  int16          `gorm:"not null"`
	Sunday    int16          `gorm:"not null"`
	StartDate datatypes.Date `gorm:"not null"`
	EndDate   datatypes.Date `gorm:"not null"`
	//Trip      Trip            `gorm:"foreignKey:ServiceId"`
	//UniversalCalendar UniversalCalendar `gorm:"foreignKey:ServiceId"`
}

func (Calendar) TableName() string {
	return "calendar"
}

func (c Calendar) GetServiceId() any {
	return c.ServiceId
}

func (c Calendar) GetMonday() any {
	return c.Monday
}

func (c Calendar) GetTuesday() any {
	return c.Tuesday
}

func (c Calendar) GetWednesday() any {
	return c.Wednesday
}

func (c Calendar) GetThursday() any {
	return c.Thursday
}

func (c Calendar) GetFriday() any {
	return c.Friday
}

func (c Calendar) GetSaturday() any {
	return c.Saturday
}

func (c Calendar) GetSunday() any {
	return c.Sunday
}

func (c Calendar) GetStartDate() any {
	return c.StartDate
}

func (c Calendar) GetEndDate() any {
	return c.EndDate
}
