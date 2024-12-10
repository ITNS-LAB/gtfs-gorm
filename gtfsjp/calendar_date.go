package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/csvutil"
	"gorm.io/datatypes"
)

type CalendarDate struct {
	ServiceId     string         `gorm:"primaryKey"`
	Date          datatypes.Date `gorm:"primaryKey"`
	ExceptionType int            `gorm:"not null"`
}

func (CalendarDate) TableName() string {
	return "calendar_dates"
}

func ParseCalendarDates(path string) ([]CalendarDate, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open calendar_dates CSV: %w", err)
	}

	// データを解析して CalendarDate 構造体のスライスを作成
	var calendarDates []CalendarDate
	for i := 0; i < len(df.Records); i++ {
		serviceId, err := df.GetString(i, "service_id")
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

		// CalendarDate 構造体を作成しリストに追加
		calendarDates = append(calendarDates, CalendarDate{
			ServiceId:     serviceId,
			Date:          date,
			ExceptionType: exceptionType,
		})
	}

	return calendarDates, nil
}
