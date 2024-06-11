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
	StopSequence      int64                  `gorm:"primaryKey;index;not null"`
	StopHeadsign      sql.NullString
	PickupType        sql.NullInt16 `gorm:"default:0"`
	DropOffType       sql.NullInt16 `gorm:"default:0"`
	ContinuousPickup  sql.NullInt16 `gorm:"default:1"`
	ContinuousDropOff sql.NullInt16 `gorm:"default:1"`
	ShapeDistTraveled sql.NullFloat64
	Timepoint         sql.NullInt16 `gorm:"default:1"`
	Trip              Trip          `gorm:"foreignKey:TripId"`
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
	if st.PickupType.Valid {
		return st.PickupType.Int16
	}
	return nil
}

func (st StopTime) GetDropOffType() any {
	if st.DropOffType.Valid {
		return st.DropOffType.Int16
	}
	return nil
}

func (st StopTime) GetContinuousPickup() any {
	if st.ContinuousPickup.Valid {
		return st.ContinuousPickup.Int16
	}
	return nil
}

func (st StopTime) GetContinuousDropOff() any {
	if st.ContinuousDropOff.Valid {
		return st.ContinuousDropOff.Int16
	}
	return nil
}

func (st StopTime) GetShapeDistTraveled() any {
	if st.ShapeDistTraveled.Valid {
		return st.ShapeDistTraveled.Float64
	}
	return nil
}

func (st StopTime) GetTimepoint() any {
	if st.Timepoint.Valid {
		return st.Timepoint.Int16
	}
	return nil
}
