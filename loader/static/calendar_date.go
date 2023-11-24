package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func LoadCalendarDates(path string) []static.CalendarDate {
	var calendarDates []static.CalendarDate
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		calendarDates = append(calendarDates, static.CalendarDate{
			ServiceId:     dataframe.IsBlank(df.GetElement("service_id")),
			Date:          dataframe.ParseDate(df.GetElement("date")),
			ExceptionType: dataframe.ParseEnum(df.GetElement("exception_type")),
		})
	}
	return calendarDates
}
