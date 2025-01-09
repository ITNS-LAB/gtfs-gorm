package gtfsschedule

import "gorm.io/datatypes"

type StopTimes struct {
	TripId                   string         `gorm:"primary_key"`
	ArrivalTime              datatypes.Time `gorm:"index;not null"`
	DepartureTime            datatypes.Time `gorm:"index;not null"`
	StopId                   string
	LocationGroupId          *string
	LocationId               *string
	StopSequence             int `gorm:"primaryKey"`
	StopHeadsign             *string
	StartPickupDropOffWindow *datatypes.Time
	EndPickupDropOffWindow   *datatypes.Time
	PickupType               *int
	DropOffType              *int
	ContinuousPickup         *int
	ContinuousDropOff        *int
	ShapeDistTraveled        *float64
	Timepoint                *int
	PickupBookingRuleId      *string
	DropOffBookingRuleid     *string
}
