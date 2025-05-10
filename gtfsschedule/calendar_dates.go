package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type CalendarDates struct {
	ServiceId     string         `gorm:"primaryKey"`
	Date          datatypes.Date `gorm:"primaryKey"`
	ExceptionType int            `gorm:"not null"`
}

func (CalendarDates) TableName() string {
	return "calendar_dates"
}

func ParseCalendarDates(path string) ([]CalendarDates, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open calendar_dates CSV: %w", err)
	}

	// データを解析して CalendarDates 構造体のスライスを作成
	var calendarDates []CalendarDates
	for i := 0; i < len(df.Records); i++ {
		serviceID, err := df.GetString(i, "service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'service_id' at row %d: %w", i, err)
		}

		date, err := df.GetDate(i, "date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'date' at row %d: %w", i, err)
		}

		exceptionType, err := df.GetInt(i, "exception_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'exception_type' at row %d: %w", i, err)
		}

		// CalendarDates 構造体を作成しリストに追加
		calendarDates = append(calendarDates, CalendarDates{
			ServiceId:     serviceID,
			Date:          date,
			ExceptionType: exceptionType,
		})
	}

	return calendarDates, nil
}

type CalendarDatesGeom struct {
	ServiceId     string         `gorm:"primaryKey"`
	Date          datatypes.Date `gorm:"primaryKey"`
	ExceptionType int            `gorm:"not null"`
}

func (CalendarDatesGeom) TableName() string {
	return "calendar_dates"
}

func ParseCalendarDatesGeom(path string) ([]CalendarDatesGeom, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open calendar_dates CSV: %w", err)
	}

	// データを解析して CalendarDates 構造体のスライスを作成
	var calendarDates []CalendarDatesGeom
	for i := 0; i < len(df.Records); i++ {
		serviceID, err := df.GetString(i, "service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'service_id' at row %d: %w", i, err)
		}

		date, err := df.GetDate(i, "date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'date' at row %d: %w", i, err)
		}

		exceptionType, err := df.GetInt(i, "exception_type")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'exception_type' at row %d: %w", i, err)
		}

		// CalendarDates 構造体を作成しリストに追加
		calendarDates = append(calendarDates, CalendarDatesGeom{
			ServiceId:     serviceID,
			Date:          date,
			ExceptionType: exceptionType,
		})
	}

	return calendarDates, nil
}
