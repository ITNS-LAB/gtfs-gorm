package realtime

import "time"

type VehiclePosition struct {
	VehiclePositionId   uint                             `gorm:"primaryKey;auto_increment"`
	Trip                VehiclePositionTripDescriptor    `gorm:"foreignKey:VehiclePositionId"`
	Vehicle             VehiclePositionVehicleDescriptor `gorm:"foreignKey:VehiclePositionId"`
	Position            Position                         `gorm:"foreignKey:VehiclePositionId"`
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
	SerialId             uint `gorm:"primaryKey;auto_increment"`
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
	SerialId          uint `gorm:"primaryKey;auto_increment"`
	VehiclePositionId uint
	Id                *string
	Label             *string
	LicensePlate      *string
}

func (VehiclePositionVehicleDescriptor) TableName() string {
	return "vehicle_position.vehicle"
}

type Position struct {
	SerialId          uint `gorm:"primaryKey;auto_increment"`
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
