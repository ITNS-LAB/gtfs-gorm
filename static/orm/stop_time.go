package orm

import "gorm.io/datatypes"

type StopTime struct {
	TripId            *string         `gorm:"primaryKey;index;not null" gorm:"index" gorm:"not null"`
	ArrivalTime       *datatypes.Time `gorm:"index"`
	DepartureTime     *datatypes.Time `gorm:"index"`
	StopId            *string         `gorm:"primaryKey;index;not null" gorm:"index" gorm:"not null"`
	StopSequence      *int            `gorm:"primaryKey;index;not null" gorm:"index" gorm:"not null"`
	StopHeadsign      *string
	PickupType        *int `gorm:"default:0"`
	DropOffType       *int `gorm:"default:0"`
	ContinuousPickup  *int `gorm:"default:1"`
	ContinuousDropOff *int `gorm:"default:1"`
	ShapeDistTraveled *float64
	Timepoint         *int `gorm:"default:1"`
	Trip              Trip `gorm:"foreignKey:TripId"`
}

func (StopTime) TableName() string {
	return "stop_times"
}
