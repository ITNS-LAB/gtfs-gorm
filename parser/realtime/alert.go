package realtime

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/realtime"
	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
	"os"
)

func AlertPbSlice(file string) []realtime.Alert {
	// gtfs-realtime ファイルをバイナリ形式で読み込む
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ファイルの読み込みエラー:", err)
	}

	// データをデシリアライズする
	feed := &gtfs.FeedMessage{}
	if err := proto.Unmarshal(data, feed); err != nil {
		fmt.Println("デシリアライズエラー:", err)
	}

	var res []realtime.Alert

	for _, entity := range feed.Entity {
		var activePeriods []realtime.TimeRange
		var informedEntity []realtime.EntitySelector
		var cause *string
		var effect *string
		var url []realtime.UrlTranslation
		var headerText []realtime.HeaderTextTranslation
		var description []realtime.DescriptionTranslation

		for _, ap := range entity.Alert.ActivePeriod {
			var start *uint64
			var end *uint64
			start = ap.Start
			end = ap.End
			activePeriods = append(activePeriods, realtime.TimeRange{
				Start: start,
				End:   end,
			})
		}

		for _, ie := range entity.Alert.InformedEntity {
			var agencyId *string
			var routeId *string
			var routeType *int32
			var directionId *uint32
			var trip realtime.AlertTripDescriptor
			var stopId *string
			agencyId = ie.AgencyId
			routeId = ie.RouteId
			routeType = ie.RouteType
			directionId = ie.DirectionId
			if ie.Trip != nil {
				trip = castTripDescriptor(ie.Trip)
			}
			stopId = ie.StopId

			informedEntity = append(informedEntity, realtime.EntitySelector{
				AgencyId:    agencyId,
				RouteId:     routeId,
				RouteType:   routeType,
				DirectionId: directionId,
				Trip:        trip,
				StopId:      stopId,
			})
		}

		if entity.Alert.Cause == nil {
			cause = nil
		} else {
			tmp := entity.Alert.GetCause().String()
			cause = &tmp
		}

		if entity.Alert.Effect == nil {
			effect = nil
		} else {
			tmp := entity.Alert.GetEffect().String()
			effect = &tmp
		}

		for _, ut := range entity.Alert.Url.Translation {
			var text *string
			var language *string
			text = ut.Text
			language = ut.Language
			url = append(url, realtime.UrlTranslation{Text: text, Language: language})
		}

		for _, ht := range entity.Alert.HeaderText.Translation {
			var text *string
			var language *string
			text = ht.Text
			language = ht.Language
			headerText = append(headerText, realtime.HeaderTextTranslation{Text: text, Language: language})
		}

		for _, dt := range entity.Alert.DescriptionText.Translation {
			var text *string
			var language *string
			text = dt.Text
			language = dt.Language
			description = append(description, realtime.DescriptionTranslation{Text: text, Language: language})
		}

		res = append(res, realtime.Alert{
			ActivePeriod:   activePeriods,
			InformedEntity: informedEntity,
			Cause:          cause,
			Effect:         effect,
			Url:            url,
			HeaderText:     headerText,
			Description:    description,
		})
	}
	return res
}

// pbからgormの構造体にキャスト
func castTripDescriptor(td *gtfs.TripDescriptor) realtime.AlertTripDescriptor {
	var tripId *string
	var routeId *string
	var directionId *uint32
	var startTime *string
	var startDate *string
	var scheduleRelationship *string
	tripId = td.TripId
	routeId = td.RouteId
	directionId = td.DirectionId
	startTime = td.StartTime
	startDate = td.StartDate
	if td.ScheduleRelationship == nil {
		scheduleRelationship = nil
	} else {
		tmp := td.GetScheduleRelationship().String()
		scheduleRelationship = &tmp
	}

	atd := realtime.AlertTripDescriptor{
		TripId:               tripId,
		RouteId:              routeId,
		DirectionId:          directionId,
		StartTime:            startTime,
		StartDate:            startDate,
		ScheduleRelationship: scheduleRelationship,
	}
	return atd
}
