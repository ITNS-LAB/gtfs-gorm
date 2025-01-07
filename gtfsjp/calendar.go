package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type Calendar struct {
	ServiceId string         `gorm:"primaryKey;index"`
	Monday    int            `gorm:"not null"`
	Tuesday   int            `gorm:"not null"`
	Wednesday int            `gorm:"not null"`
	Thursday  int            `gorm:"not null"`
	Friday    int            `gorm:"not null"`
	Saturday  int            `gorm:"not null"`
	Sunday    int            `gorm:"not null"`
	StartDate datatypes.Date `gorm:"not null"`
	EndDate   datatypes.Date `gorm:"not null"`
	//Trips     []Trip         `gorm:"foreignKey:ServiceId;references:ServiceId"` calendar.txtが存在しない場合に、外部キー制約に引っかかるためコメントアウト
}

func (Calendar) TableName() string {
	return "calendar"
}

func ParseCalendar(path string) ([]Calendar, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open calendar CSV: %w", err)
	}

	// データを解析して Calendar 構造体のスライスを作成
	var calendars []Calendar
	for i := 0; i < len(df.Records); i++ {
		serviceId, err := df.GetString(i, "service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'service_id' at row %d: %w", i, err)
		}

		monday, err := df.GetInt(i, "monday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'monday' at row %d: %w", i, err)
		}

		tuesday, err := df.GetInt(i, "tuesday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'tuesday' at row %d: %w", i, err)
		}

		wednesday, err := df.GetInt(i, "wednesday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'wednesday' at row %d: %w", i, err)
		}

		thursday, err := df.GetInt(i, "thursday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'thursday' at row %d: %w", i, err)
		}

		friday, err := df.GetInt(i, "friday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'friday' at row %d: %w", i, err)
		}

		saturday, err := df.GetInt(i, "saturday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'saturday' at row %d: %w", i, err)
		}

		sunday, err := df.GetInt(i, "sunday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'sunday' at row %d: %w", i, err)
		}

		startDate, err := df.GetDate(i, "start_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'start_date' at row %d: %w", i, err)
		}

		endDate, err := df.GetDate(i, "end_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'end_date' at row %d: %w", i, err)
		}

		// Calendar 構造体を作成しリストに追加
		calendars = append(calendars, Calendar{
			ServiceId: serviceId,
			Monday:    monday,
			Tuesday:   tuesday,
			Wednesday: wednesday,
			Thursday:  thursday,
			Friday:    friday,
			Saturday:  saturday,
			Sunday:    sunday,
			StartDate: startDate,
			EndDate:   endDate,
		})
	}

	return calendars, nil
}

type CalendarGeom struct {
	ServiceId string         `gorm:"primaryKey;index"`
	Monday    int            `gorm:"not null"`
	Tuesday   int            `gorm:"not null"`
	Wednesday int            `gorm:"not null"`
	Thursday  int            `gorm:"not null"`
	Friday    int            `gorm:"not null"`
	Saturday  int            `gorm:"not null"`
	Sunday    int            `gorm:"not null"`
	StartDate datatypes.Date `gorm:"not null"`
	EndDate   datatypes.Date `gorm:"not null"`
	//Trips     []TripGeom     `gorm:"foreignKey:ServiceId;references:ServiceId"` calendar.txtが存在しない場合に、外部キー制約に引っかかるためコメントアウト
}

func (CalendarGeom) TableName() string {
	return "calendar"
}

func ParseCalendarGeom(path string) ([]CalendarGeom, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open calendar CSV: %w", err)
	}

	// データを解析して Calendar 構造体のスライスを作成
	var calendars []CalendarGeom
	for i := 0; i < len(df.Records); i++ {
		serviceId, err := df.GetString(i, "service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'service_id' at row %d: %w", i, err)
		}

		monday, err := df.GetInt(i, "monday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'monday' at row %d: %w", i, err)
		}

		tuesday, err := df.GetInt(i, "tuesday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'tuesday' at row %d: %w", i, err)
		}

		wednesday, err := df.GetInt(i, "wednesday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'wednesday' at row %d: %w", i, err)
		}

		thursday, err := df.GetInt(i, "thursday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'thursday' at row %d: %w", i, err)
		}

		friday, err := df.GetInt(i, "friday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'friday' at row %d: %w", i, err)
		}

		saturday, err := df.GetInt(i, "saturday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'saturday' at row %d: %w", i, err)
		}

		sunday, err := df.GetInt(i, "sunday")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'sunday' at row %d: %w", i, err)
		}

		startDate, err := df.GetDate(i, "start_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'start_date' at row %d: %w", i, err)
		}

		endDate, err := df.GetDate(i, "end_date")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'end_date' at row %d: %w", i, err)
		}

		// Calendar 構造体を作成しリストに追加
		calendars = append(calendars, CalendarGeom{
			ServiceId: serviceId,
			Monday:    monday,
			Tuesday:   tuesday,
			Wednesday: wednesday,
			Thursday:  thursday,
			Friday:    friday,
			Saturday:  saturday,
			Sunday:    sunday,
			StartDate: startDate,
			EndDate:   endDate,
		})
	}

	return calendars, nil
}
