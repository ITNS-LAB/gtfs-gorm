package parsestatic

import (
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
			return []ormstatic.CalendarDate{}, err
		}

		serviceId, err := dataframe.ParseString(df.GetElement("service_id"))
		if err != nil {
			return []ormstatic.CalendarDate{}, err
		}

		date, err := dataframe.ParseDataTypesDate(df.GetElement("date"))
		if err != nil {
			return []ormstatic.CalendarDate{}, err
		}

		exceptionType, err := dataframe.ParseInt16(df.GetElement("exception_type"))
		if err != nil {
			return []ormstatic.CalendarDate{}, err
		}

		calendarDates = append(calendarDates, ormstatic.CalendarDate{
			ServiceId:     serviceId,
			Date:          date,
			ExceptionType: exceptionType,
		})
	}
	return calendarDates, nil
}
