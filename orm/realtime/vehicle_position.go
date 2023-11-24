package realtime

import "time"

type VehiclePosition struct {
	VehiclePositionId   uint                             `gorm:"primaryKey" gorm:"auto_increment"`
	Trip                VehiclePositionTripDescriptor    `gorm:"foreignkey:VehiclePositionId"`
	Vehicle             VehiclePositionVehicleDescriptor `gorm:"foreignkey:VehiclePositionId"`
	Position            Position                         `gorm:"foreignkey:VehiclePositionId"`
	CurrentStopSequence *uint32
	StopId              *string
	CurrentStatus       *string
	TimeStamp           *uint64
	DateTime            *time.Time
	CongestionLevel     *string
	OccupancyStatus     *string
}

func (VehiclePosition) TableName() string {
	return "vehicle_position.vehicle_position"
}

type VehiclePositionTripDescriptor struct {
	SerialId             uint `gorm:"primaryKey" gorm:"auto_increment"`
	VehiclePositionId    uint
	TripId               *string
	RouteId              *string
	DirectionId          *uint32
	StartTime            *string
	StartDate            *string
	ScheduleRelationship *string
}

func (VehiclePositionTripDescriptor) TableName() string {
	return "vehicle_position.trip"
}

type VehiclePositionVehicleDescriptor struct {
	SerialId          uint `gorm:"primaryKey" gorm:"auto_increment"`
	VehiclePositionId uint
	Id                *string
	Label             *string
	LicensePlate      *string
}

func (VehiclePositionVehicleDescriptor) TableName() string {
	return "vehicle_position.vehicle"
}

type Position struct {
	SerialId          uint `gorm:"primaryKey" gorm:"auto_increment"`
	VehiclePositionId uint
	Latitude          *float32
	Longitude         *float32
	Bearing           *float32
	Odometer          *float64
	Speed             *float32
}

func (Position) TableName() string {
	return "vehicle_position.position"
}
