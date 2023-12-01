package ormstatic

import "gorm.io/datatypes"

type Frequency struct {
	TripId      *string         `gorm:"primaryKey;index;not null"`
	StartTime   *datatypes.Time `gorm:"index;not null"`
	EndTime     *datatypes.Time `gorm:"index;not null"`
	HeadwaySecs *int            `gorm:"not null"`
	ExactTimes  *int            `gorm:"default:0"`
}

func (Frequency) TableName() string {
	return "frequencies"
}
