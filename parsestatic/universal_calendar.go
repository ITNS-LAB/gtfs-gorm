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
		week[0] = calendar.GetSunday().(int)
		week[1] = calendar.GetMonday().(int)
		week[2] = calendar.GetTuesday().(int)
		week[3] = calendar.GetWednesday().(int)
		week[4] = calendar.GetThursday().(int)
		week[5] = calendar.GetFriday().(int)
		week[6] = calendar.GetSaturday().(int)

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
					Date:      date,
				}
			}
			_ = ucMap
		}
	}

	// calendar_dateの中身を確認し、例外を処理
	for _, i := range calendarDates {
		if i.ExceptionType == 1 {
			date, _ := i.Date.Value()
			ucMap[date.(time.Time)] = ormstatic.UniversalCalendar{
				ServiceId: i.ServiceId,
				Date:      i.Date,
			}
			continue
		}
		if i.ExceptionType == 2 {
			date, _ := i.Date.Value()
			delete(ucMap, date.(time.Time))
			continue
		}
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
