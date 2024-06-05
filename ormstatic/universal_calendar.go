package ormstatic

import "gorm.io/datatypes"

type UniversalCalendar struct {
	ServiceId string         `gorm:"primaryKey;index"`
	Date      datatypes.Date `gorm:"primaryKey;index"`
}

func (UniversalCalendar) TableName() string {
	return "universal_calendar"
}

func (uc UniversalCalendar) GetServiceId() any {
	return uc.ServiceId
}

func (uc UniversalCalendar) GetDate() any {
	return uc.Date
}
