package ormstatic

import (
	"database/sql"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/nulldatatypes"
)

type StopTime struct {
	TripId            string                 `gorm:"primaryKey;index;not null"`
	ArrivalTime       nulldatatypes.NullTime `gorm:"index"`
	DepartureTime     nulldatatypes.NullTime `gorm:"index"`
	StopId            string                 `gorm:"primaryKey;index;not null"`
	StopSequence      int                    `gorm:"primaryKey;index;not null"`
	StopHeadsign      sql.NullString
	PickupType        int16 `gorm:"default:0"`
	DropOffType       int16 `gorm:"default:0"`
	ContinuousPickup  int16 `gorm:"default:1"`
	ContinuousDropOff int16 `gorm:"default:1"`
	ShapeDistTraveled sql.NullFloat64
	Timepoint         int16 `gorm:"default:1"`
	Trip              Trip  `gorm:"foreignKey:TripId"`
}

func (StopTime) TableName() string {
	return "stop_times"
}

func (st StopTime) GetTripId() any {
	return st.TripId
}

func (st StopTime) GetArrivalTime() any {
	if st.ArrivalTime.Valid {
		return st.ArrivalTime.Time
	}
	return nil
}

func (st StopTime) GetDepartureTime() any {
	if st.DepartureTime.Valid {
		return st.DepartureTime.Time
	}
	return nil
}

func (st StopTime) GetStopId() any {
	return st.StopId
}

func (st StopTime) GetStopSequence() any {
	return st.StopSequence
}

func (st StopTime) GetStopHeadsign() any {
	if st.StopHeadsign.Valid {
		return st.StopHeadsign.String
	}
	return nil
}

func (st StopTime) GetPickupType() any {
	return st.PickupType
}

func (st StopTime) GetDropOffType() any {
	return st.DropOffType
}

func (st StopTime) GetContinuousPickup() any {
	return st.ContinuousPickup
}

func (st StopTime) GetContinuousDropOff() any {
	return st.ContinuousDropOff
}

func (st StopTime) GetShapeDistTraveled() any {
	if st.ShapeDistTraveled.Valid {
		return st.ShapeDistTraveled.Float64
	}
	return nil
}

func (st StopTime) GetTimepoint() any {
	return st.Timepoint
}
