package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseCalendarDates(path string) []orm.CalendarDate {
	var calendarDates []orm.CalendarDate
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		calendarDates = append(calendarDates, orm.CalendarDate{
			ServiceId:     dataframe.IsBlank(df.GetElement("service_id")),
			Date:          dataframe.ParseDate(df.GetElement("date")),
			ExceptionType: dataframe.ParseEnum(df.GetElement("exception_type")),
		})
	}
	return calendarDates
}
