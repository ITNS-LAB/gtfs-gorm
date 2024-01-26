package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseCalendarDates(path string) ([]ormstatic.CalendarDate, error) {
	var calendarDates []ormstatic.CalendarDate
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return calendarDates, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		calendarDates = append(calendarDates, ormstatic.CalendarDate{
			ServiceId:     dataframe.IsBlank(df.GetElement("service_id")),
			Date:          dataframe.ParseDate(df.GetElement("date")),
			ExceptionType: dataframe.ParseEnum(df.GetElement("exception_type")),
		})
	}
	return calendarDates, nil
}
