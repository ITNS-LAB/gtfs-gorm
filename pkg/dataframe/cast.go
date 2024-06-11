package dataframe

import (
	"database/sql"
	"errors"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/nulldatatypes"
	"gorm.io/datatypes"
	"strconv"
	"time"
)

func ParseString(ns sql.NullString) (string, error) {
	if ns.Valid {
		return ns.String, nil
	}
	return "", errors.New("invalid string")
}

func ParseInt16(ns sql.NullString) (int16, error) {
	if ns.Valid {
		i, err := strconv.Atoi(ns.String)
		if err != nil {
			return 0, errors.New("not a valid integer")
		}
		if i < -32768 || i > 32767 {
			return 0, errors.New("out of int16 range")
		}
		return int16(i), nil
	}
	return 0, errors.New("invalid string")
}

func ParseNullInt16(ns sql.NullString) (sql.NullInt16, error) {
	if ns.Valid {
		i, err := strconv.Atoi(ns.String)
		if err != nil {
			return sql.NullInt16{Int16: 0, Valid: false}, errors.New("not a valid integer")
		}
		if i < -32768 || i > 32767 {
			return sql.NullInt16{Int16: 0, Valid: false}, errors.New("out of int16 range")
		}
		return sql.NullInt16{Int16: int16(i), Valid: true}, nil
	}
	return sql.NullInt16{Int16: 0, Valid: false}, nil
}

func ParseInt32(ns sql.NullString) (int32, error) {
	if ns.Valid {
		i, err := strconv.Atoi(ns.String)
		if err != nil {
			return 0, errors.New("not a valid integer")
		}
		if i < -2147483648 || i > 2147483647 {
			return 0, errors.New("out of int32 range")
		}
		return int32(i), nil
	}
	return 0, errors.New("invalid string")
}

func ParseNullInt32(ns sql.NullString) (sql.NullInt32, error) {
	if ns.Valid {
		i, err := strconv.Atoi(ns.String)
		if err != nil {
			return sql.NullInt32{Int32: 0, Valid: false}, errors.New("not a valid integer")
		}
		if i < -2147483648 || i > 2147483647 {
			return sql.NullInt32{Int32: 0, Valid: false}, errors.New("out of int32 range")
		}
		return sql.NullInt32{Int32: int32(i), Valid: true}, nil
	}
	return sql.NullInt32{Int32: 0, Valid: false}, nil
}

func ParseInt64(ns sql.NullString) (int64, error) {
	if ns.Valid {
		i, err := strconv.Atoi(ns.String)
		if err != nil {
			return 0, errors.New("not a valid integer")
		}
		return int64(i), nil
	}
	return 0, errors.New("invalid string")
}

func ParseNullInt64(ns sql.NullString) (sql.NullInt64, error) {
	if ns.Valid {
		i, err := strconv.Atoi(ns.String)
		if err != nil {
			return sql.NullInt64{Int64: 0, Valid: false}, errors.New("not a valid integer")
		}
		return sql.NullInt64{Int64: int64(i), Valid: true}, nil
	}
	return sql.NullInt64{Int64: 0, Valid: false}, nil
}

func ParseFloat64(ns sql.NullString) (float64, error) {
	if ns.Valid {
		f, err := strconv.ParseFloat(ns.String, 64)
		if err != nil {
			return 0, errors.New("not a valid float")
		}
		return f, nil
	}
	return 0, errors.New("invalid string")
}

func ParseNullFloat64(ns sql.NullString) (sql.NullFloat64, error) {
	if ns.Valid {
		f, err := strconv.ParseFloat(ns.String, 64)
		if err != nil {
			return sql.NullFloat64{Float64: 0, Valid: false}, errors.New("not a valid float")
		}
		return sql.NullFloat64{Float64: f, Valid: true}, nil
	}
	return sql.NullFloat64{Float64: 0, Valid: false}, nil
}

func ParseTime(ns sql.NullString) (time.Time, error) {
	if ns.Valid {
		layout := "20060102"
		t, err := time.Parse(layout, ns.String)
		if err != nil {
			return time.Time{}, errors.New("invalid date format")
		}
		return t, nil
	}
	return time.Time{}, errors.New("invalid string")
}

func ParseDataTypesDate(ns sql.NullString) (datatypes.Date, error) {
	if ns.Valid {
		layout := "20060102"
		t, err := time.Parse(layout, ns.String)
		if err != nil {
			return datatypes.Date{}, errors.New("invalid date format")
		}
		dtd := datatypes.Date(t)
		return dtd, nil
	}
	return datatypes.Date{}, errors.New("invalid string")
}

func ParseNullDataTypesDate(ns sql.NullString) (nulldatatypes.NullDate, error) {
	if ns.Valid {
		layout := "20060102"
		t, err := time.Parse(layout, ns.String)
		if err != nil {
			return nulldatatypes.NullDate{Date: datatypes.Date{}, Valid: false}, errors.New("invalid date format")
		}
		dtd := datatypes.Date(t)
		return nulldatatypes.NullDate{Date: dtd, Valid: true}, nil
	}
	return nulldatatypes.NullDate{Date: datatypes.Date{}, Valid: false}, nil
}

func ParseDataTypesTime(ns sql.NullString) (datatypes.Time, error) {
	if ns.Valid {
		layout := "15:04:05"
		t, err := time.Parse(layout, ns.String)
		if err != nil {
			return datatypes.NewTime(0, 0, 0, 0), errors.New("invalid time format")
		}
		dtt := datatypes.NewTime(t.Hour(), t.Minute(), t.Second(), 0)
		return dtt, nil
	}
	return datatypes.NewTime(0, 0, 0, 0), errors.New("invalid string")
}

func ParseNullDataTypesTime(ns sql.NullString) (nulldatatypes.NullTime, error) {
	if ns.Valid {
		layout := "15:04:05"
		t, err := time.Parse(layout, ns.String)
		if err != nil {
			return nulldatatypes.NullTime{Time: datatypes.NewTime(0, 0, 0, 0), Valid: false},
				errors.New("invalid time format")
		}
		dtt := datatypes.NewTime(t.Hour(), t.Minute(), t.Second(), 0)
		return nulldatatypes.NullTime{Time: dtt, Valid: true}, nil
	}
	return nulldatatypes.NullTime{Time: datatypes.NewTime(0, 0, 0, 0), Valid: false}, nil
}
