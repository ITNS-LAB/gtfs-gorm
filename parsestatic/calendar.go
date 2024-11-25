package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseCalendar(path string) ([]gtfsjp.Calendar, error) {
	var calendars []gtfsjp.Calendar
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

		calendars = append(calendars, gtfsjp.Calendar{
			ServiceId: util.IsBlank(df.GetElement("service_id")),
			Monday:    util.ParseEnum(df.GetElement("monday")),
			Tuesday:   util.ParseEnum(df.GetElement("tuesday")),
			Wednesday: util.ParseEnum(df.GetElement("wednesday")),
			Thursday:  util.ParseEnum(df.GetElement("thursday")),
			Friday:    util.ParseEnum(df.GetElement("friday")),
			Saturday:  util.ParseEnum(df.GetElement("saturday")),
			Sunday:    util.ParseEnum(df.GetElement("sunday")),
			StartDate: util.ParseDate(df.GetElement("start_date")),
			EndDate:   util.ParseDate(df.GetElement("end_date")),
		})
	}
	return calendars, nil
}
