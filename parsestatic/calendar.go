package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseCalendar(path string) ([]ormstatic.Calendar, error) {
	var calendars []ormstatic.Calendar
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return calendars, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			return calendars, err
		}

		serviceId, err := dataframe.ParseString(df.GetElement("service_id"))
		monday, err := dataframe.ParseInt16(df.GetElement("monday"))
		tuesday, err := dataframe.ParseInt16(df.GetElement("tuesday"))
		wednesday, err := dataframe.ParseInt16(df.GetElement("wednesday"))
		thursday, err := dataframe.ParseInt16(df.GetElement("thursday"))
		friday, err := dataframe.ParseInt16(df.GetElement("friday"))
		saturday, err := dataframe.ParseInt16(df.GetElement("saturday"))
		sunday, err := dataframe.ParseInt16(df.GetElement("sunday"))
		startdate, err := dataframe.ParseDataTypesDate(df.GetElement("start_date"))
		endDate, err := dataframe.ParseDataTypesDate(df.GetElement("end_date"))

		calendars = append(calendars, ormstatic.Calendar{
			ServiceId: serviceId,
			Monday:    monday,
			Tuesday:   tuesday,
			Wednesday: wednesday,
			Thursday:  thursday,
			Friday:    friday,
			Saturday:  saturday,
			Sunday:    sunday,
			StartDate: startdate,
			EndDate:   endDate,
		})
	}
	return calendars, nil
}
