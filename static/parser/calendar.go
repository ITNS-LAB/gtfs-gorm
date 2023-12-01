package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseCalendar(path string) []orm.Calendar {
	var calendars []orm.Calendar
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		calendars = append(calendars, orm.Calendar{
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
	return calendars
}
