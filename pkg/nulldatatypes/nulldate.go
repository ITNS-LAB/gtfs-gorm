package nulldatatypes

import (
	"database/sql/driver"
	"gorm.io/datatypes"
)

type NullDate struct {
	datatypes.Date
	Valid bool
}

func (NullDate) GormDataType() string {
	return "date"
}

func (n *NullDate) Scan(value any) error {
	if value == nil {
		n.Date, n.Valid = datatypes.Date{}, false
		return nil
	}
	n.Valid = true
	return n.Date.Scan(value)
}

func (n NullDate) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Date.Value()
}

//func (n NullDate) Get() any {
//	if n.Valid {
//		return n.Date
//	}
//	return nil
//}
