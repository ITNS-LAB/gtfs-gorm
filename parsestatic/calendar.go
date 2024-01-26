package parsestatic

import (
	"fmt"
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
			fmt.Println("Error:", err)
			break
		}

		calendars = append(calendars, ormstatic.Calendar{
			ServiceId: dataframe.IsBlank(df.GetElement("service_id")),
			Monday:    dataframe.ParseEnum(df.GetElement("monday")),
			Tuesday:   dataframe.ParseEnum(df.GetElement("tuesday")),
			Wednesday: dataframe.ParseEnum(df.GetElement("wednesday")),
			Thursday:  dataframe.ParseEnum(df.GetElement("thursday")),
			Friday:    dataframe.ParseEnum(df.GetElement("friday")),
			Saturday:  dataframe.ParseEnum(df.GetElement("saturday")),
			Sunday:    dataframe.ParseEnum(df.GetElement("sunday")),
			StartDate: dataframe.ParseDate(df.GetElement("start_date")),
			EndDate:   dataframe.ParseDate(df.GetElement("end_date")),
		})
	}
	return calendars, nil
}
