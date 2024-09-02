package parserealtime

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormrealtime"
	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
	"os"
)

func TripUpdatePbSlice(file string) ([]ormrealtime.TripUpdate, error) {
	var res []ormrealtime.TripUpdate

	// gtfs.zip-ormrealtime ファイルをバイナリ形式で読み込む
	data, err := os.ReadFile(file)
	if err != nil {
		return res, err
	}

	// データをデシリアライズする
	feed := &gtfs.FeedMessage{}
	if err := proto.Unmarshal(data, feed); err != nil {
		return res, err
	}

	for _, entity := range feed.Entity {
		var trip ormrealtime.TripUpdateTripDescriptor
		var vehicle ormrealtime.TripUpdateVehicleDescriptor
		var stopTimeUpdate []ormrealtime.StopTimeUpdate
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
			var arrival ormrealtime.ArrivalStopTimeEvent
			var departure ormrealtime.DepartureStopTimeEvent
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

			stopTimeUpdate = append(stopTimeUpdate, ormrealtime.StopTimeUpdate{
				StopSequence:         stopSequence,
				StopId:               stopId,
				Arrival:              arrival,
				Departure:            departure,
				ScheduleRelationship: scheduleRelationship,
			})
		}

		timeStamp = entity.TripUpdate.Timestamp
		delay = entity.TripUpdate.Delay

		res = append(res, ormrealtime.TripUpdate{
			Trip:           trip,
			Vehicle:        vehicle,
			StopTimeUpdate: stopTimeUpdate,
			TimeStamp:      timeStamp,
			Delay:          delay,
		})

	}
	return res, nil
}

func TripUpdatePbMap(file string) (map[string]ormrealtime.TripUpdate, error) {
	// mapの初期化
	res := make(map[string]ormrealtime.TripUpdate)

	// gtfs.zip-ormrealtime ファイルをバイナリ形式で読み込む
	data, err := os.ReadFile(file)
	if err != nil {
		return res, err
	}

	// データをデシリアライズする
	feed := &gtfs.FeedMessage{}
	if err := proto.Unmarshal(data, feed); err != nil {
		return res, err
	}

	for _, entity := range feed.Entity {
		var trip ormrealtime.TripUpdateTripDescriptor
		var vehicle ormrealtime.TripUpdateVehicleDescriptor
		var stopTimeUpdate []ormrealtime.StopTimeUpdate
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
			var arrival ormrealtime.ArrivalStopTimeEvent
			var departure ormrealtime.DepartureStopTimeEvent
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

			stopTimeUpdate = append(stopTimeUpdate, ormrealtime.StopTimeUpdate{
				StopSequence:         stopSequence,
				StopId:               stopId,
				Arrival:              arrival,
				Departure:            departure,
				ScheduleRelationship: scheduleRelationship,
			})
		}

		timeStamp = entity.TripUpdate.Timestamp
		delay = entity.TripUpdate.Delay

		res[*trip.TripId] = ormrealtime.TripUpdate{
			Trip:           trip,
			Vehicle:        vehicle,
			StopTimeUpdate: stopTimeUpdate,
			TimeStamp:      timeStamp,
			Delay:          delay,
		}
	}
	return res, nil
}
