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
			return []ormstatic.Calendar{}, err
		}

		serviceId, err := dataframe.ParseString(df.GetElement("service_id"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		monday, err := dataframe.ParseInt16(df.GetElement("monday"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		tuesday, err := dataframe.ParseInt16(df.GetElement("tuesday"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		wednesday, err := dataframe.ParseInt16(df.GetElement("wednesday"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		thursday, err := dataframe.ParseInt16(df.GetElement("thursday"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		friday, err := dataframe.ParseInt16(df.GetElement("friday"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		saturday, err := dataframe.ParseInt16(df.GetElement("saturday"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		sunday, err := dataframe.ParseInt16(df.GetElement("sunday"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		startDate, err := dataframe.ParseDataTypesDate(df.GetElement("start_date"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		endDate, err := dataframe.ParseDataTypesDate(df.GetElement("end_date"))
		if err != nil {
			return []ormstatic.Calendar{}, err
		}

		calendars = append(calendars, ormstatic.Calendar{
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
