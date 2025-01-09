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
	StartPickupDropOffWindow *string
	EndPickupDropOffWindow   *string
	PickupType               *string
	DropOffType              *string
	ContinuousPickup         *string
	ContinuousDropOff        *string
	ShapeDistTraveled        *string
	Timepoint                *string
	PickupBookingRuleId      *string
	DropOffBookingRuleid     *string
}
