package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
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
			ServiceId:     util.IsBlank(df.GetElement("service_id")),
			Date:          util.ParseDate(df.GetElement("date")),
			ExceptionType: util.ParseEnum(df.GetElement("exception_type")),
		})
	}
	return calendarDates, nil
}
