package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/realtime/orm"
	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
	"os"
)

func TripUpdatePbSlice(file string) []orm.TripUpdate {
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

	var res []orm.TripUpdate

	for _, entity := range feed.Entity {
		var trip orm.TripUpdateTripDescriptor
		var vehicle orm.TripUpdateVehicleDescriptor
		var stopTimeUpdate []orm.StopTimeUpdate
		var timeStamp *uint64
		var delay *int32

		if entity.TripUpdate.Trip != nil {
			trip.TripId = entity.TripUpdate.Trip.TripId
			trip.RouteId = entity.TripUpdate.Trip.RouteId
			trip.DirectionId = entity.TripUpdate.Trip.DirectionId
			trip.StartTime = entity.TripUpdate.Trip.StartTime
			trip.StartDate = entity.TripUpdate.Trip.StartDate
			if entity.TripUpdate.Trip.ScheduleRelationship == nil {
				trip.ScheduleRelationship = nil
			} else {
				tmp := entity.TripUpdate.Trip.GetScheduleRelationship().String()
				trip.ScheduleRelationship = &tmp
			}
		}

		if entity.TripUpdate.Vehicle != nil {
			vehicle.Id = entity.TripUpdate.Vehicle.Id
			vehicle.Label = entity.TripUpdate.Vehicle.Label
			vehicle.LicensePlate = entity.TripUpdate.Vehicle.LicensePlate
		}

		for _, stu := range entity.TripUpdate.StopTimeUpdate {
			var stopSequence *uint32
			var stopId *string
			var arrival orm.ArrivalStopTimeEvent
			var departure orm.DepartureStopTimeEvent
			var scheduleRelationship *string

			stopSequence = stu.StopSequence
			stopId = stu.StopId
			if stu.Arrival != nil {
				arrival.Delay = stu.Arrival.Delay
				arrival.Time = stu.Arrival.Time
				arrival.Uncertainty = stu.Arrival.Uncertainty
			}
			if stu.Departure != nil {
				departure.Delay = stu.Departure.Delay
				departure.Time = stu.Departure.Time
				departure.Uncertainty = stu.Departure.Uncertainty
			}
			if stu.ScheduleRelationship == nil {
				scheduleRelationship = nil
			} else {
				tmp := stu.GetScheduleRelationship().String()
				scheduleRelationship = &tmp
			}

			stopTimeUpdate = append(stopTimeUpdate, orm.StopTimeUpdate{
				StopSequence:         stopSequence,
				StopId:               stopId,
				Arrival:              arrival,
				Departure:            departure,
				ScheduleRelationship: scheduleRelationship,
			})
		}

		timeStamp = entity.TripUpdate.Timestamp
		delay = entity.TripUpdate.Delay

		res = append(res, orm.TripUpdate{
			Trip:           trip,
			Vehicle:        vehicle,
			StopTimeUpdate: stopTimeUpdate,
			TimeStamp:      timeStamp,
			Delay:          delay,
		})

	}
	return res
}

func TripUpdatePbMap(file string) map[string]orm.TripUpdate {
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

	// mapの初期化
	res := make(map[string]orm.TripUpdate)

	for _, entity := range feed.Entity {
		var trip orm.TripUpdateTripDescriptor
		var vehicle orm.TripUpdateVehicleDescriptor
		var stopTimeUpdate []orm.StopTimeUpdate
		var timeStamp *uint64
		var delay *int32

		if entity.TripUpdate.Trip != nil {
			trip.TripId = entity.TripUpdate.Trip.TripId
			trip.RouteId = entity.TripUpdate.Trip.RouteId
			trip.DirectionId = entity.TripUpdate.Trip.DirectionId
			trip.StartTime = entity.TripUpdate.Trip.StartTime
			trip.StartDate = entity.TripUpdate.Trip.StartDate
			if entity.TripUpdate.Trip.ScheduleRelationship == nil {
				trip.ScheduleRelationship = nil
			} else {
				tmp := entity.TripUpdate.Trip.GetScheduleRelationship().String()
				trip.ScheduleRelationship = &tmp
			}
		}

		if entity.TripUpdate.Vehicle != nil {
			vehicle.Id = entity.TripUpdate.Vehicle.Id
			vehicle.Label = entity.TripUpdate.Vehicle.Label
			vehicle.LicensePlate = entity.TripUpdate.Vehicle.LicensePlate
		}

		for _, stu := range entity.TripUpdate.StopTimeUpdate {
			var stopSequence *uint32
			var stopId *string
			var arrival orm.ArrivalStopTimeEvent
			var departure orm.DepartureStopTimeEvent
			var scheduleRelationship *string

			stopSequence = stu.StopSequence
			stopId = stu.StopId
			if stu.Arrival != nil {
				arrival.Delay = stu.Arrival.Delay
				arrival.Time = stu.Arrival.Time
				arrival.Uncertainty = stu.Arrival.Uncertainty
			}
			if stu.Departure == nil {
				departure.Delay = stu.Departure.Delay
				departure.Time = stu.Departure.Time
				departure.Uncertainty = stu.Departure.Uncertainty
			}
			if stu.ScheduleRelationship == nil {
				scheduleRelationship = nil
			} else {
				tmp := stu.GetScheduleRelationship().String()
				scheduleRelationship = &tmp
			}

			stopTimeUpdate = append(stopTimeUpdate, orm.StopTimeUpdate{
				StopSequence:         stopSequence,
				StopId:               stopId,
				Arrival:              arrival,
				Departure:            departure,
				ScheduleRelationship: scheduleRelationship,
			})
		}

		timeStamp = entity.TripUpdate.Timestamp
		delay = entity.TripUpdate.Delay

		res[*trip.TripId] = orm.TripUpdate{
			Trip:           trip,
			Vehicle:        vehicle,
			StopTimeUpdate: stopTimeUpdate,
			TimeStamp:      timeStamp,
			Delay:          delay,
		}
	}
	return res
}
