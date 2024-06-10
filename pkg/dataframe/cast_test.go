package dataframe

import (
	"database/sql"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/nulldatatypes"
	"gorm.io/datatypes"
	"reflect"
	"testing"
	"time"
)

func TestParseString(t *testing.T) {
	type args struct {
		ns sql.NullString
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid string", args{ns: sql.NullString{String: "valid string", Valid: true}}, "valid string", false},
		{"null string", args{ns: sql.NullString{String: "", Valid: false}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseString(tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt16(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    int16
		wantErr bool
	}{
		{"valid int16", sql.NullString{String: "12345", Valid: true}, 12345, false},
		{"invalid int16", sql.NullString{String: "abc", Valid: true}, 0, true},
		{"out of range int16", sql.NullString{String: "40000", Valid: true}, 0, true},
		{"null string", sql.NullString{String: "", Valid: false}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt16(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNullInt16(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    sql.NullInt16
		wantErr bool
	}{
		{"valid int16", sql.NullString{String: "12345", Valid: true}, sql.NullInt16{Int16: 12345, Valid: true}, false},
		{"invalid int16", sql.NullString{String: "abc", Valid: true}, sql.NullInt16{Int16: 0, Valid: false}, true},
		{"out of range int16", sql.NullString{String: "40000", Valid: true}, sql.NullInt16{Int16: 0, Valid: false}, true},
		{"null string", sql.NullString{String: "", Valid: false}, sql.NullInt16{Int16: 0, Valid: false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNullInt16(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNullInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNullInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt32(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    int32
		wantErr bool
	}{
		{"valid int32", sql.NullString{String: "123456789", Valid: true}, 123456789, false},
		{"invalid int32", sql.NullString{String: "abc", Valid: true}, 0, true},
		{"out of range int32", sql.NullString{String: "4000000000", Valid: true}, 0, true},
		{"null string", sql.NullString{String: "", Valid: false}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt32(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNullInt32(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    sql.NullInt32
		wantErr bool
	}{
		{"valid int32", sql.NullString{String: "123456789", Valid: true}, sql.NullInt32{Int32: 123456789, Valid: true}, false},
		{"invalid int32", sql.NullString{String: "abc", Valid: true}, sql.NullInt32{Int32: 0, Valid: false}, true},
		{"out of range int32", sql.NullString{String: "4000000000", Valid: true}, sql.NullInt32{Int32: 0, Valid: false}, true},
		{"null string", sql.NullString{String: "", Valid: false}, sql.NullInt32{Int32: 0, Valid: false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNullInt32(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNullInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNullInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt64(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    int64
		wantErr bool
	}{
		{"valid int64", sql.NullString{String: "1234567890123", Valid: true}, 1234567890123, false},
		{"invalid int64", sql.NullString{String: "abc", Valid: true}, 0, true},
		{"null string", sql.NullString{String: "", Valid: false}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt64(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNullInt64(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    sql.NullInt64
		wantErr bool
	}{
		{"valid int64", sql.NullString{String: "1234567890123", Valid: true}, sql.NullInt64{Int64: 1234567890123, Valid: true}, false},
		{"invalid int64", sql.NullString{String: "abc", Valid: true}, sql.NullInt64{Int64: 0, Valid: false}, true},
		{"null string", sql.NullString{String: "", Valid: false}, sql.NullInt64{Int64: 0, Valid: false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNullInt64(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNullInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseNullInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    time.Time
		wantErr bool
	}{
		{"valid date string", sql.NullString{String: "20230611", Valid: true}, time.Date(2023, 6, 11, 0, 0, 0, 0, time.UTC), false},
		{"invalid date string", sql.NullString{String: "invalid", Valid: true}, time.Time{}, true},
		{"null string", sql.NullString{String: "", Valid: false}, time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTime(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDataTypesDate(t *testing.T) {
	type args struct {
		ns sql.NullString
	}
	tests := []struct {
		name    string
		args    args
		want    datatypes.Date
		wantErr bool
	}{
		{"valid date string", args{ns: sql.NullString{String: "20230611", Valid: true}}, datatypes.Date(time.Date(2023, 6, 11, 0, 0, 0, 0, time.UTC)), false},
		{"invalid date string", args{ns: sql.NullString{String: "invalid", Valid: true}}, datatypes.Date{}, true},
		{"null string", args{ns: sql.NullString{String: "", Valid: false}}, datatypes.Date{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDataTypesDate(tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDataTypesDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDataTypesDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNullDataTypesDate(t *testing.T) {
	type args struct {
		ns sql.NullString
	}
	tests := []struct {
		name    string
		args    args
		want    nulldatatypes.NullDate
		wantErr bool
	}{
		{"valid date string", args{ns: sql.NullString{String: "20230611", Valid: true}}, nulldatatypes.NullDate{Date: datatypes.Date(time.Date(2023, 6, 11, 0, 0, 0, 0, time.UTC)), Valid: true}, false},
		{"invalid date string", args{ns: sql.NullString{String: "invalid", Valid: true}}, nulldatatypes.NullDate{Date: datatypes.Date{}, Valid: false}, true},
		{"null string", args{ns: sql.NullString{String: "", Valid: false}}, nulldatatypes.NullDate{Date: datatypes.Date{}, Valid: false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNullDataTypesDate(tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNullDataTypesDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNullDataTypesDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDataTypesTime(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    datatypes.Time
		wantErr bool
	}{
		{"valid time string", sql.NullString{String: "15:04:05", Valid: true}, datatypes.NewTime(15, 4, 5, 0), false},
		{"invalid time string", sql.NullString{String: "invalid", Valid: true}, datatypes.NewTime(0, 0, 0, 0), true},
		{"null string", sql.NullString{String: "", Valid: false}, datatypes.NewTime(0, 0, 0, 0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDataTypesTime(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDataTypesTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDataTypesTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNullDataTypesTime(t *testing.T) {
	tests := []struct {
		name    string
		ns      sql.NullString
		want    nulldatatypes.NullTime
		wantErr bool
	}{
		{"valid time string", sql.NullString{String: "15:04:05", Valid: true}, nulldatatypes.NullTime{Time: datatypes.NewTime(15, 4, 5, 0), Valid: true}, false},
		{"invalid time string", sql.NullString{String: "invalid", Valid: true}, nulldatatypes.NullTime{Time: datatypes.NewTime(0, 0, 0, 0), Valid: false}, true},
		{"null string", sql.NullString{String: "", Valid: false}, nulldatatypes.NullTime{Time: datatypes.NewTime(0, 0, 0, 0), Valid: false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNullDataTypesTime(tt.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNullDataTypesTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNullDataTypesTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
