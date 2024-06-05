package ormstatic

import "database/sql"

type Transfer struct {
	Id              int    `gorm:"primaryKey;auto_increment;not null"`
	FromStopId      string `gorm:"primaryKey;not null"`
	ToStopId        string `gorm:"primaryKey;not null"`
	FromRouteId     sql.NullString
	ToRouteId       sql.NullString
	FromTripId      sql.NullString
	ToTripId        sql.NullString
	TransferType    int `gorm:"not null"`
	MinTransferTime sql.NullInt32
	FromStop        Stop `gorm:"foreignKey:FromStopId;references:StopId"`
	ToStop          Stop `gorm:"foreignKey:ToStopId;references:StopId"`
}

func (Transfer) TableName() string {
	return "transfers"
}

func (t Transfer) GetId() any {
	return t.Id
}

func (t Transfer) GetFromStopId() any {
	return t.FromStopId
}

func (t Transfer) GetToStopId() any {
	return t.ToStopId
}

func (t Transfer) GetFromRouteId() any {
	if t.FromRouteId.Valid {
		return t.FromRouteId.String
	}
	return nil
}

func (t Transfer) GetToRouteId() any {
	if t.ToRouteId.Valid {
		return t.ToRouteId.String
	}
	return nil
}

func (t Transfer) GetFromTripId() any {
	if t.FromTripId.Valid {
		return t.FromTripId.String
	}
	return nil
}

func (t Transfer) GetToTripId() any {
	if t.ToTripId.Valid {
		return t.ToTripId.String
	}
	return nil
}

func (t Transfer) GetTransferType() any {
	return t.TransferType
}

func (t Transfer) GetMinTransferTime() any {
	if t.MinTransferTime.Valid {
		return t.MinTransferTime.Int32
	}
	return nil
}
