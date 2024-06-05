package nulldatatypes

import (
	"database/sql/driver"
	"gorm.io/datatypes"
)

type NullTime struct {
	datatypes.Time
	Valid bool
}

func (NullTime) GormDataType() string {
	return "time"
}

func (nt NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}
	nt.Valid = true
	return nt.Time.Scan(value)
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time.Value()
}
