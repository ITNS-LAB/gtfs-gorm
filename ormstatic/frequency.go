package ormstatic

import "gorm.io/datatypes"

type Frequency struct {
	TripId      *string         `gorm:"primaryKey"`
	StartTime   *datatypes.Time `gorm:"primaryKey"`
	EndTime     *datatypes.Time `gorm:"primaryKey"`
	HeadwaySecs *int            `gorm:"not null"`
	ExactTimes  *int            `gorm:"default:0"`
}

func (Frequency) TableName() string {
	return "frequencies"
}
