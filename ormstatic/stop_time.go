package ormstatic

import "gorm.io/datatypes"

type StopTime struct {
	TripId            *string         `gorm:"primaryKey"`
	ArrivalTime       *datatypes.Time `gorm:"index"`
	DepartureTime     *datatypes.Time `gorm:"index"`
	StopId            *string         `gorm:"primaryKey"`
	StopSequence      *int            `gorm:"primaryKey"`
	StopHeadsign      *string
	PickupType        *int `gorm:"default:0"`
	DropOffType       *int `gorm:"default:0"`
	ContinuousPickup  *int `gorm:"default:1"`
	ContinuousDropOff *int `gorm:"default:1"`
	ShapeDistTraveled *float64
	Timepoint         *int `gorm:"default:1"`
}

func (StopTime) TableName() string {
	return "stop_times"
}
