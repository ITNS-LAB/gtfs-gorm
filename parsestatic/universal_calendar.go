package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"gorm.io/datatypes"
	"sort"
	"time"
)

func ParseUniversalCalendar(calendars []ormstatic.Calendar, calendarDates []ormstatic.CalendarDate) []ormstatic.UniversalCalendar {
	var uc []ormstatic.UniversalCalendar

	// universalCalendarのマップ初期化
	ucMap := make(map[time.Time]ormstatic.UniversalCalendar)

	// calendarの内容を展開しucMapに挿入
	for _, calendar := range calendars {
		week := make(map[time.Weekday]int)
		week[0] = *calendar.Sunday
		week[1] = *calendar.Monday
		week[2] = *calendar.Tuesday
		week[3] = *calendar.Wednesday
		week[4] = *calendar.Thursday
		week[5] = *calendar.Friday
		week[6] = *calendar.Saturday

		s, _ := calendar.StartDate.Value()
		e, _ := calendar.EndDate.Value()
		start := s.(time.Time)
		end := e.(time.Time)
		// start_dateからend_dateまでのループ
		for t := start; t.Before(end); t = t.AddDate(0, 0, 1) {
			//fmt.Println(t)
			//fmt.Println(t.Weekday())
			date := datatypes.Date(t)
			if week[t.Weekday()] == 1 {
				ucMap[t] = ormstatic.UniversalCalendar{
					ServiceId: calendar.ServiceId,
					Date:      &date,
				}
			}
			_ = ucMap
		}
	}

	// calendar_dateの中身を確認し、例外を処理
	for _, i := range calendarDates {
		if i.ExceptionType == nil {
			panic("'exceptionType' is nil")
		}
		if *i.ExceptionType == 0 {
			date, _ := i.Date.Value()
			delete(ucMap, date.(time.Time))
			continue
		}
		if *i.ExceptionType == 1 {
			date, _ := i.Date.Value()
			ucMap[date.(time.Time)] = ormstatic.UniversalCalendar{
				ServiceId: i.ServiceId,
				Date:      i.Date,
			}
			continue
		}
		panic("'exception type' is not a valid value")
	}

	// ucMapのキーを取り出し
	var keys []time.Time
	for k := range ucMap {
		keys = append(keys, k)
	}
	// キーのスライスをソート
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})
	// ソートされたキーを使ってmapの要素にアクセス
	for _, k := range keys {
		uc = append(uc, ucMap[k])
	}
	return uc
}
