package ormrealtime

import (
	"database/sql"
	"time"
)

type VehiclePosition struct {
	VehiclePositionId   uint                             `gorm:"primaryKey;auto_increment"`
	Trip                VehiclePositionTripDescriptor    `gorm:"foreignKey:VehiclePositionId"`
	Vehicle             VehiclePositionVehicleDescriptor `gorm:"foreignKey:VehiclePositionId"`
	Position            Position                         `gorm:"foreignKey:VehiclePositionId"`
	CurrentStopSequence sql.NullInt32
	StopId              sql.NullString
	CurrentStatus       sql.NullString
	TimeStamp           sql.NullInt64
	DateTime            *time.Time
	CongestionLevel     sql.NullString
	OccupancyStatus     sql.NullString
}

func (VehiclePosition) TableName() string {
	return "vehicle_position.vehicle_position"
}

func (v VehiclePosition) GetVehiclePositionId() any {
	return v.VehiclePositionId
}

func (v VehiclePosition) GetTrip() any {
	return v.Trip
}

func (v VehiclePosition) GetVehicle() any {
	return v.Vehicle
}

func (v VehiclePosition) GetPosition() any {
	return v.Position
}

func (v VehiclePosition) GetCurrentStopSequence() any {
	if v.CurrentStopSequence.Valid {
		return v.CurrentStopSequence.Int32
	}
	return nil
}

func (v VehiclePosition) GetStopId() any {
	if v.StopId.Valid {
		return v.StopId.String
	}
	return nil
}

func (v VehiclePosition) GetCurrentStatus() any {
	if v.CurrentStatus.Valid {
		return v.CurrentStatus.String
	}
	return nil
}

func (v VehiclePosition) GetTimeStamp() any {
	if v.TimeStamp.Valid {
		return v.TimeStamp.Int64
	}
	return nil
}

func (v VehiclePosition) GetDateTime() any {
	return v.DateTime
}

func (v VehiclePosition) GetCongestionLevel() any {
	if v.CongestionLevel.Valid {
		return v.CongestionLevel.String
	}
	return nil
}

func (v VehiclePosition) GetOccupancyStatus() any {
	if v.OccupancyStatus.Valid {
		return v.OccupancyStatus.String
	}
	return nil
}

type VehiclePositionTripDescriptor struct {
	SerialId             uint `gorm:"primaryKey;auto_increment"`
	VehiclePositionId    uint
	TripId               sql.NullString
	RouteId              sql.NullString
	DirectionId          sql.NullInt32
	StartTime            sql.NullString
	StartDate            sql.NullString
	ScheduleRelationship sql.NullString
}

func (VehiclePositionTripDescriptor) TableName() string {
	return "vehicle_position.trip"
}

func (v VehiclePositionTripDescriptor) GetSerialId() any {
	return v.SerialId
}

func (v VehiclePositionTripDescriptor) GetVehiclePositionId() any {
	return v.VehiclePositionId
}

func (v VehiclePositionTripDescriptor) GetTripId() any {
	if v.TripId.Valid {
		return v.TripId.String
	}
	return nil
}

func (v VehiclePositionTripDescriptor) GetRouteId() any {
	if v.RouteId.Valid {
		return v.RouteId.String
	}
	return nil
}

func (v VehiclePositionTripDescriptor) GetDirectionId() any {
	if v.DirectionId.Valid {
		return v.DirectionId.Int32
	}
	return nil
}

func (v VehiclePositionTripDescriptor) GetStartTime() any {
	if v.StartTime.Valid {
		return v.StartTime.String
	}
	return nil
}

func (v VehiclePositionTripDescriptor) GetStartDate() any {
	if v.StartDate.Valid {
		return v.StartDate.String
	}
	return nil
}

func (v VehiclePositionTripDescriptor) GetScheduleRelationship() any {
	if v.ScheduleRelationship.Valid {
		return v.ScheduleRelationship.String
	}
	return nil
}

type VehiclePositionVehicleDescriptor struct {
	SerialId          uint `gorm:"primaryKey;auto_increment"`
	VehiclePositionId uint
	Id                sql.NullString
	Label             sql.NullString
	LicensePlate      sql.NullString
}

func (VehiclePositionVehicleDescriptor) TableName() string {
	return "vehicle_position.vehicle"
}

func (v VehiclePositionVehicleDescriptor) GetSerialId() any {
	return v.SerialId
}

func (v VehiclePositionVehicleDescriptor) GetVehiclePositionId() any {
	return v.VehiclePositionId
}

func (v VehiclePositionVehicleDescriptor) GetId() any {
	if v.Id.Valid {
		return v.Id.String
	}
	return nil
}

func (v VehiclePositionVehicleDescriptor) GetLabel() any {
	if v.Label.Valid {
		return v.Label.String
	}
	return nil
}

func (v VehiclePositionVehicleDescriptor) GetLicensePlate() any {
	if v.LicensePlate.Valid {
		return v.LicensePlate.String
	}
	return nil
}

type Position struct {
	SerialId          uint `gorm:"primaryKey;auto_increment"`
	VehiclePositionId uint
	Latitude          sql.NullFloat64
	Longitude         sql.NullFloat64
	Bearing           sql.NullFloat64
	Odometer          sql.NullFloat64
	Speed             sql.NullFloat64
}

func (Position) TableName() string {
	return "vehicle_position.position"
}

func (p Position) GetSerialId() any {
	return p.SerialId
}

func (p Position) GetVehiclePositionId() any {
	return p.VehiclePositionId
}

func (p Position) GetLatitude() any {
	if p.Latitude.Valid {
		return p.Latitude.Float64
	}
	return nil
}

func (p Position) GetLongitude() any {
	if p.Longitude.Valid {
		return p.Longitude.Float64
	}
	return nil
}

func (p Position) GetBearing() any {
	if p.Bearing.Valid {
		return p.Bearing.Float64
	}
	return nil
}

func (p Position) GetOdometer() any {
	if p.Odometer.Valid {
		return p.Odometer.Float64
	}
	return nil
}

func (p Position) GetSpeed() any {
	if p.Speed.Valid {
		return p.Speed.Float64
	}
	return nil
}
