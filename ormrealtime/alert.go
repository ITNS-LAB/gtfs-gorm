package ormrealtime

import "database/sql"

type Alert struct {
	AlertId        uint             `gorm:"primaryKey;auto_increment"`
	ActivePeriod   []TimeRange      `gorm:"foreignKey:AlertId"`
	InformedEntity []EntitySelector `gorm:"foreignKey:AlertId"`
	Cause          sql.NullString
	Effect         sql.NullString
	Url            []UrlTranslation         `gorm:"foreignKey:AlertId"`
	HeaderText     []HeaderTextTranslation  `gorm:"foreignKey:AlertId"`
	Description    []DescriptionTranslation `gorm:"foreignKey:AlertId"`
}

func (Alert) TableName() string {
	return "alert.alert"
}

func (a Alert) GetAlertId() any {
	return a.AlertId
}

func (a Alert) GetCause() any {
	if a.Cause.Valid {
		return a.Cause.String
	}
	return nil
}

func (a Alert) GetEffect() any {
	if a.Effect.Valid {
		return a.Effect.String
	}
	return nil
}

type TimeRange struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Start    sql.NullInt64
	End      sql.NullInt64
}

func (TimeRange) TableName() string {
	return "alert.active_period"
}

func (tr TimeRange) GetSerialId() any {
	return tr.SerialId
}

func (tr TimeRange) GetAlertId() any {
	return tr.AlertId
}

func (tr TimeRange) GetStart() any {
	if tr.Start.Valid {
		return tr.Start.Int64
	}
	return nil
}

func (tr TimeRange) GetEnd() any {
	if tr.End.Valid {
		return tr.End.Int64
	}
	return nil
}

type EntitySelector struct {
	InformedEntityId uint `gorm:"primaryKey;auto_increment"`
	AlertId          uint
	AgencyId         sql.NullString
	RouteId          sql.NullString
	RouteType        sql.NullInt32
	DirectionId      sql.NullInt32
	Trip             AlertTripDescriptor `gorm:"foreignKey:InformedEntityId"`
	StopId           sql.NullString
}

func (EntitySelector) TableName() string {
	return "alert.informed_entity"
}

func (es EntitySelector) GetInformedEntityId() any {
	return es.InformedEntityId
}

func (es EntitySelector) GetAlertId() any {
	return es.AlertId
}

func (es EntitySelector) GetAgencyId() any {
	if es.AgencyId.Valid {
		return es.AgencyId.String
	}
	return nil
}

func (es EntitySelector) GetRouteId() any {
	if es.RouteId.Valid {
		return es.RouteId.String
	}
	return nil
}

func (es EntitySelector) GetRouteType() any {
	if es.RouteType.Valid {
		return es.RouteType.Int32
	}
	return nil
}

func (es EntitySelector) GetDirectionId() any {
	if es.DirectionId.Valid {
		return es.DirectionId.Int32
	}
	return nil
}

func (es EntitySelector) GetStopId() any {
	if es.StopId.Valid {
		return es.StopId.String
	}
	return nil
}

type AlertTripDescriptor struct {
	SerialId             uint `gorm:"primaryKey;auto_increment"`
	InformedEntityId     uint
	TripId               sql.NullString
	RouteId              sql.NullString
	DirectionId          sql.NullInt32
	StartTime            sql.NullString
	StartDate            sql.NullString
	ScheduleRelationship sql.NullString
}

func (AlertTripDescriptor) TableName() string {
	return "alert.trip"
}

func (atd AlertTripDescriptor) GetSerialId() any {
	return atd.SerialId
}

func (atd AlertTripDescriptor) GetInformedEntityId() any {
	return atd.InformedEntityId
}

func (atd AlertTripDescriptor) GetTripId() any {
	if atd.TripId.Valid {
		return atd.TripId.String
	}
	return nil
}

func (atd AlertTripDescriptor) GetRouteId() any {
	if atd.RouteId.Valid {
		return atd.RouteId.String
	}
	return nil
}

func (atd AlertTripDescriptor) GetDirectionId() any {
	if atd.DirectionId.Valid {
		return atd.DirectionId.Int32
	}
	return nil
}

func (atd AlertTripDescriptor) GetStartTime() any {
	if atd.StartTime.Valid {
		return atd.StartTime.String
	}
	return nil
}

func (atd AlertTripDescriptor) GetStartDate() any {
	if atd.StartDate.Valid {
		return atd.StartDate.String
	}
	return nil
}

func (atd AlertTripDescriptor) GetScheduleRelationship() any {
	if atd.ScheduleRelationship.Valid {
		return atd.ScheduleRelationship.String
	}
	return nil
}

type UrlTranslation struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Text     sql.NullString
	Language sql.NullString
}

func (UrlTranslation) TableName() string {
	return "alert.url"
}

func (ut UrlTranslation) GetSerialId() any {
	return ut.SerialId
}

func (ut UrlTranslation) GetAlertId() any {
	return ut.AlertId
}

func (ut UrlTranslation) GetText() any {
	if ut.Text.Valid {
		return ut.Text.String
	}
	return nil
}

func (ut UrlTranslation) GetLanguage() any {
	if ut.Language.Valid {
		return ut.Language.String
	}
	return nil
}

type HeaderTextTranslation struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Text     sql.NullString
	Language sql.NullString
}

func (HeaderTextTranslation) TableName() string {
	return "alert.header_text"
}

func (htt HeaderTextTranslation) GetSerialId() any {
	return htt.SerialId
}

func (htt HeaderTextTranslation) GetAlertId() any {
	return htt.AlertId
}

func (htt HeaderTextTranslation) GetText() any {
	if htt.Text.Valid {
		return htt.Text.String
	}
	return nil
}

func (htt HeaderTextTranslation) GetLanguage() any {
	if htt.Language.Valid {
		return htt.Language.String
	}
	return nil
}

type DescriptionTranslation struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Text     sql.NullString
	Language sql.NullString
}

func (DescriptionTranslation) TableName() string {
	return "alert.description_text"
}

func (dt DescriptionTranslation) GetSerialId() any {
	return dt.SerialId
}

func (dt DescriptionTranslation) GetAlertId() any {
	return dt.AlertId
}

func (dt DescriptionTranslation) GetText() any {
	if dt.Text.Valid {
		return dt.Text.String
	}
	return nil
}

func (dt DescriptionTranslation) GetLanguage() any {
	if dt.Language.Valid {
		return dt.Language.String
	}
	return nil
}
