package ormrealtime

import (
	"database/sql"
)

type TripUpdate struct {
	TripUpdateId   uint                        `gorm:"primaryKey;auto_increment"`
	Trip           TripUpdateTripDescriptor    `gorm:"foreignKey:TripUpdateId"`
	Vehicle        TripUpdateVehicleDescriptor `gorm:"foreignKey:TripUpdateId"`
	StopTimeUpdate []StopTimeUpdate            `gorm:"foreignKey:TripUpdateId"`
	TimeStamp      sql.NullInt64
	DateTime       sql.NullTime
	Delay          sql.NullInt32
}

func (TripUpdate) TableName() string {
	return "trip_update.trip_update"
}

func (t TripUpdate) GetTripUpdateId() any {
	return t.TripUpdateId
}

func (t TripUpdate) GetTrip() any {
	return t.Trip
}

func (t TripUpdate) GetVehicle() any {
	return t.Vehicle
}

func (t TripUpdate) GetTimeStamp() any {
	if t.TimeStamp.Valid {
		return t.TimeStamp.Int64
	}
	return nil
}

func (t TripUpdate) GetDateTime() any {
	if t.DateTime.Valid {
		return t.DateTime.Time
	}
	return nil
}

func (t TripUpdate) GetDelay() any {
	if t.Delay.Valid {
		return t.Delay.Int32
	}
	return nil
}

type TripUpdateTripDescriptor struct {
	SerialId             uint `gorm:"primaryKey;auto_increment"`
	TripUpdateId         uint
	TripId               sql.NullString
	RouteId              sql.NullString
	DirectionId          sql.NullInt32
	StartTime            sql.NullString
	StartDate            sql.NullString
	ScheduleRelationship sql.NullString
}

func (TripUpdateTripDescriptor) TableName() string {
	return "trip_update.trip"
}

func (t TripUpdateTripDescriptor) GetSerialId() any {
	return t.SerialId
}

func (t TripUpdateTripDescriptor) GetTripUpdateId() any {
	return t.TripUpdateId
}

func (t TripUpdateTripDescriptor) GetTripId() any {
	if t.TripId.Valid {
		return t.TripId.String
	}
	return nil
}

func (t TripUpdateTripDescriptor) GetRouteId() any {
	if t.RouteId.Valid {
		return t.RouteId.String
	}
	return nil
}

func (t TripUpdateTripDescriptor) GetDirectionId() any {
	if t.DirectionId.Valid {
		return t.DirectionId.Int32
	}
	return nil
}

func (t TripUpdateTripDescriptor) GetStartTime() any {
	if t.StartTime.Valid {
		return t.StartTime.String
	}
	return nil
}

func (t TripUpdateTripDescriptor) GetStartDate() any {
	if t.StartDate.Valid {
		return t.StartDate.String
	}
	return nil
}

func (t TripUpdateTripDescriptor) GetScheduleRelationship() any {
	if t.ScheduleRelationship.Valid {
		return t.ScheduleRelationship.String
	}
	return nil
}

type TripUpdateVehicleDescriptor struct {
	SerialId     uint `gorm:"primaryKey;auto_increment"`
	TripUpdateId uint
	Id           sql.NullString
	Label        sql.NullString
	LicensePlate sql.NullString
}

func (TripUpdateVehicleDescriptor) TableName() string {
	return "trip_update.vehicle"
}

func (t TripUpdateVehicleDescriptor) GetSerialId() any {
	return t.SerialId
}

func (t TripUpdateVehicleDescriptor) GetTripUpdateId() any {
	return t.TripUpdateId
}

func (t TripUpdateVehicleDescriptor) GetId() any {
	if t.Id.Valid {
		return t.Id.String
	}
	return nil
}

func (t TripUpdateVehicleDescriptor) GetLabel() any {
	if t.Label.Valid {
		return t.Label.String
	}
	return nil
}

func (t TripUpdateVehicleDescriptor) GetLicensePlate() any {
	if t.LicensePlate.Valid {
		return t.LicensePlate.String
	}
	return nil
}

type StopTimeUpdate struct {
	StopTimeUpdateId     uint `gorm:"primaryKey;auto_increment"`
	TripUpdateId         uint
	StopSequence         sql.NullInt32
	StopId               sql.NullString
	Arrival              ArrivalStopTimeEvent   `gorm:"foreignKey:StopTimeUpdateId"`
	Departure            DepartureStopTimeEvent `gorm:"foreignKey:StopTimeUpdateId"`
	ScheduleRelationship sql.NullString
}

func (StopTimeUpdate) TableName() string {
	return "trip_update.stop_time_update"
}

func (s StopTimeUpdate) GetStopTimeUpdateId() any {
	return s.StopTimeUpdateId
}

func (s StopTimeUpdate) GetTripUpdateId() any {
	return s.TripUpdateId
}

func (s StopTimeUpdate) GetStopSequence() any {
	if s.StopSequence.Valid {
		return s.StopSequence.Int32
	}
	return nil
}

func (s StopTimeUpdate) GetStopId() any {
	if s.StopId.Valid {
		return s.StopId.String
	}
	return nil
}

func (s StopTimeUpdate) GetArrival() any {
	return s.Arrival
}

func (s StopTimeUpdate) GetDeparture() any {
	return s.Departure
}

func (s StopTimeUpdate) GetScheduleRelationship() any {
	if s.ScheduleRelationship.Valid {
		return s.ScheduleRelationship.String
	}
	return nil
}

type ArrivalStopTimeEvent struct {
	SerialId         uint `gorm:"primaryKey;auto_increment"`
	StopTimeUpdateId uint
	Delay            sql.NullInt32
	Time             sql.NullInt64
	Uncertainty      sql.NullInt32
}

func (ArrivalStopTimeEvent) TableName() string {
	return "trip_update.arrival"
}

func (a ArrivalStopTimeEvent) GetSerialId() any {
	return a.SerialId
}

func (a ArrivalStopTimeEvent) GetStopTimeUpdateId() any {
	return a.StopTimeUpdateId
}

func (a ArrivalStopTimeEvent) GetDelay() any {
	if a.Delay.Valid {
		return a.Delay.Int32
	}
	return nil
}

func (a ArrivalStopTimeEvent) GetTime() any {
	if a.Time.Valid {
		return a.Time.Int64
	}
	return nil
}

func (a ArrivalStopTimeEvent) GetUncertainty() any {
	if a.Uncertainty.Valid {
		return a.Uncertainty.Int32
	}
	return nil
}

type DepartureStopTimeEvent struct {
	SerialId         uint `gorm:"primaryKey;auto_increment"`
	StopTimeUpdateId uint
	Delay            sql.NullInt32
	Time             sql.NullInt64
	Uncertainty      sql.NullInt32
}

func (DepartureStopTimeEvent) TableName() string {
	return "trip_update.departure"
}

func (d DepartureStopTimeEvent) GetSerialId() any {
	return d.SerialId
}

func (d DepartureStopTimeEvent) GetStopTimeUpdateId() any {
	return d.StopTimeUpdateId
}

func (d DepartureStopTimeEvent) GetDelay() any {
	if d.Delay.Valid {
		return d.Delay.Int32
	}
	return nil
}

func (d DepartureStopTimeEvent) GetTime() any {
	if d.Time.Valid {
		return d.Time.Int64
	}
	return nil
}

func (d DepartureStopTimeEvent) GetUncertainty() any {
	if d.Uncertainty.Valid {
		return d.Uncertainty.Int32
	}
	return nil
}
