package ormstatic

import "database/sql"

type FareRule struct {
	Id              int    `gorm:"primaryKey;auto_increment;not null"`
	FareId          string `gorm:"index;not null"`
	RouteId         sql.NullString
	OriginId        sql.NullString
	DestinationId   sql.NullString
	ContainsId      sql.NullString
	OriginStop      Stop `gorm:"foreignKey:OriginId;references:ZoneId"`
	DestinationStop Stop `gorm:"foreignKey:DestinationId;references:ZoneId"`
}

func (FareRule) TableName() string {
	return "fare_rules"
}

func (f FareRule) GetId() any {
	return f.Id
}

func (f FareRule) GetFareId() any {
	return f.FareId
}

func (f FareRule) GetRouteId() any {
	if f.RouteId.Valid {
		return f.RouteId.String
	}
	return nil
}

func (f FareRule) GetOriginId() any {
	if f.OriginId.Valid {
		return f.OriginId.String
	}
	return nil
}

func (f FareRule) GetDestinationId() any {
	if f.DestinationId.Valid {
		return f.DestinationId.String
	}
	return nil
}

func (f FareRule) GetContainsId() any {
	if f.ContainsId.Valid {
		return f.ContainsId.String
	}
	return nil
}
