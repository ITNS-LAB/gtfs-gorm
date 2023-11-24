package dataframe

import (
	"gorm.io/datatypes"
	"strconv"
	"time"
)

func IsBlank(s *string) *string {
	if s == nil {
		return nil
	}
	if *s == "" {
		return nil
	}
	return s
}

func ParseFloat64(s *string) *float64 {
	if s == nil || *s == "" {
		return nil
	}
	f, _ := strconv.ParseFloat(*s, 64)
	return &f
}

func ParseInt(s *string) *int {
	if s == nil || *s == "" {
		return nil
	}
	i, _ := strconv.Atoi(*s)
	return &i
}

// ParseEnum ParseInt()と同じ
func ParseEnum(s *string) *int {
	return ParseInt(s)
}

func ParseDate(s *string) *datatypes.Date {
	if s == nil || *s == "" {
		return nil
	}
	layout := "20060102"
	t, _ := time.Parse(layout, *s)
	dtd := datatypes.Date(t)
	return &dtd
}

func ParseTime(s *string) *datatypes.Time {
	if s == nil || *s == "" {
		return nil
	}
	layout := "15:04:05"
	t, _ := time.Parse(layout, *s)
	dtt := datatypes.NewTime(t.Hour(), t.Minute(), t.Second(), 0)
	return &dtt
}
