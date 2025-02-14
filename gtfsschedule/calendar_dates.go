package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type CalendarDates struct {
	ServiceId     string         `gorm:"primary_key"`
	Date          datatypes.Date `gorm:"not null"`
	ExceptionType int            `gorm:"not null"`
	Trips         []Trips        `gorm:"foreignKey:ServiceId;references:ServiceId"`
	TimeFrame     []TimeFrame    `gorm:"foreignKey:ServiceId;references:ServiceId"`
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
