package ormstatic

import "database/sql"

type Stop struct {
	StopId             string `gorm:"primaryKey;index;not null"`
	StopCode           sql.NullString
	StopName           sql.NullString //jpでは必須
	StopDesc           sql.NullString
	StopLat            sql.NullFloat64 //jpでは必須
	StopLon            sql.NullFloat64 //jpでは必須
	ZoneId             sql.NullString  `gorm:"unique"`
	StopUrl            sql.NullString
	LocationType       int16 `gorm:"default:0"`
	ParentStation      sql.NullString
	StopTimezone       sql.NullString
	WheelchairBoarding int16 `gorm:"default:0"`
	LevelId            sql.NullString
	PlatformCode       sql.NullString
	StopTime           StopTime `gorm:"foreignKey:StopId"`
}

func (Stop) TableName() string {
	return "stops"
}

func (s Stop) GetStopId() any {
	return s.StopId
}

func (s Stop) GetStopCode() any {
	if s.StopCode.Valid {
		return s.StopCode.String
	}
	return nil
}

func (s Stop) GetStopName() any {
	if s.StopName.Valid {
		return s.StopName.String
	}
	return nil
}

func (s Stop) GetStopDesc() any {
	if s.StopDesc.Valid {
		return s.StopDesc.String
	}
	return nil
}

func (s Stop) GetStopLat() any {
	if s.StopLat.Valid {
		return s.StopLat.Float64
	}
	return nil
}

func (s Stop) GetStopLon() any {
	if s.StopLon.Valid {
		return s.StopLon.Float64
	}
	return nil
}

func (s Stop) GetZoneId() any {
	if s.ZoneId.Valid {
		return s.ZoneId.String
	}
	return nil
}

func (s Stop) GetStopUrl() any {
	if s.StopUrl.Valid {
		return s.StopUrl.String
	}
	return nil
}

func (s Stop) GetLocationType() any {
	return s.LocationType
}

func (s Stop) GetParentStation() any {
	if s.ParentStation.Valid {
		return s.ParentStation.String
	}
	return nil
}

func (s Stop) GetStopTimezone() any {
	if s.StopTimezone.Valid {
		return s.StopTimezone.String
	}
	return nil
}

func (s Stop) GetWheelchairBoarding() any {
	return s.WheelchairBoarding
}

func (s Stop) GetLevelId() any {
	if s.LevelId.Valid {
		return s.LevelId.String
	}
	return nil
}

func (s Stop) GetPlatformCode() any {
	if s.PlatformCode.Valid {
		return s.PlatformCode.String
	}
	return nil
}

func (s Stop) GetStopTime() any {
	return s.StopTime
}
