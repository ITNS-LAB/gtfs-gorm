package realtime

import "time"

type TripUpdate struct {
	TripUpdateId   uint                        `gorm:"primaryKey" gorm:"auto_increment"`
	Trip           TripUpdateTripDescriptor    `gorm:"foreignKey:TripUpdateId"`
	Vehicle        TripUpdateVehicleDescriptor `gorm:"foreignKey:TripUpdateId"`
	StopTimeUpdate []StopTimeUpdate            `gorm:"foreignKey:TripUpdateId"`
	TimeStamp      *uint64
	DateTime       *time.Time
	Delay          *int32
}

func (TripUpdate) TableName() string {
	return "trip_update.trip_update"
}

type TripUpdateTripDescriptor struct {
	SerialId             uint `gorm:"primaryKey" gorm:"auto_increment"`
	TripUpdateId         uint
	TripId               *string
	RouteId              *string
	DirectionId          *uint32
	StartTime            *string
	StartDate            *string
	ScheduleRelationship *string
}

func (TripUpdateTripDescriptor) TableName() string {
	return "trip_update.trip"
}

type TripUpdateVehicleDescriptor struct {
	SerialId     uint `gorm:"primaryKey" gorm:"auto_increment"`
	TripUpdateId uint
	Id           *string
	Label        *string
	LicensePlate *string
}

func (TripUpdateVehicleDescriptor) TableName() string {
	return "trip_update.vehicle"
}

type StopTimeUpdate struct {
	StopTimeUpdateId     uint `gorm:"primaryKey" gorm:"auto_increment"`
	TripUpdateId         uint
	StopSequence         *uint32
	StopId               *string
	Arrival              ArrivalStopTimeEvent   `gorm:"foreignKey:StopTimeUpdateId"`
	Departure            DepartureStopTimeEvent `gorm:"foreignKey:StopTimeUpdateId"`
	ScheduleRelationship *string
}

func (StopTimeUpdate) TableName() string {
	return "trip_update.stop_time_update"
}

type ArrivalStopTimeEvent struct {
	SerialId         uint `gorm:"primaryKey" gorm:"auto_increment"`
	StopTimeUpdateId uint
	Delay            *int32
	Time             *int64
	Uncertainty      *int32
}

func (ArrivalStopTimeEvent) TableName() string {
	return "trip_update.arrival"
}

type DepartureStopTimeEvent struct {
	SerialId         uint `gorm:"primaryKey" gorm:"auto_increment"`
	StopTimeUpdateId uint
	Delay            *int32
	Time             *int64
	Uncertainty      *int32
}

func (DepartureStopTimeEvent) TableName() string {
	return "trip_update.departure"
}
