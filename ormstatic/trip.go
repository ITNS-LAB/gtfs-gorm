package ormstatic

import "database/sql"

type Trip struct {
	RouteId              string `gorm:"index;not null"`
	ServiceId            string `gorm:"index;not null"`
	TripId               string `gorm:"primaryKey;index;not null"`
	TripHeadsign         sql.NullString
	TripShortName        sql.NullString
	DirectionId          sql.NullInt16  `gorm:"index"`
	BlockId              sql.NullString `gorm:"index"`
	ShapeId              sql.NullString `gorm:"index"`
	WheelchairAccessible int16          `gorm:"default:0"`
	BikesAllowed         int16          `gorm:"default:0"`
	Frequency            Frequency      `gorm:"foreignKey:TripId"`
}

func (Trip) TableName() string {
	return "trips"
}

func (t Trip) GetRouteId() any {
	return t.RouteId
}

func (t Trip) GetServiceId() any {
	return t.ServiceId
}

func (t Trip) GetTripId() any {
	return t.TripId
}

func (t Trip) GetTripHeadsign() any {
	if t.TripHeadsign.Valid {
		return t.TripHeadsign.String
	}
	return nil
}

func (t Trip) GetTripShortName() any {
	if t.TripShortName.Valid {
		return t.TripShortName.String
	}
	return nil
}

func (t Trip) GetDirectionId() any {
	if t.DirectionId.Valid {
		return t.DirectionId.Int16
	}
	return nil
}

func (t Trip) GetBlockId() any {
	if t.BlockId.Valid {
		return t.BlockId.String
	}
	return nil
}

func (t Trip) GetShapeId() any {
	if t.ShapeId.Valid {
		return t.ShapeId.String
	}
	return nil
}

func (t Trip) GetWheelchairAccessible() any {
	return t.WheelchairAccessible
}

func (t Trip) GetBikesAllowed() any {
	return t.BikesAllowed
}
