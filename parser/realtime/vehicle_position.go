package realtime

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/realtime"
	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
	"os"
)

func VehiclePositionPbSlice(file string) []realtime.VehiclePosition {
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

	var res []realtime.VehiclePosition

	for _, entity := range feed.Entity {
		var trip realtime.VehiclePositionTripDescriptor
		var vehicle realtime.VehiclePositionVehicleDescriptor
		var position realtime.Position
		var currentStopSequence *uint32
		var stopId *string
		var currentStatus *string
		var timeStamp *uint64
		var congestionLevel *string
		var occupancyStatus *string

		if entity.Vehicle.Trip != nil {
			trip.TripId = entity.Vehicle.Trip.TripId
			trip.RouteId = entity.Vehicle.Trip.RouteId
			trip.DirectionId = entity.Vehicle.Trip.DirectionId
			trip.StartTime = entity.Vehicle.Trip.StartTime
			trip.StartDate = entity.Vehicle.Trip.StartDate
			if entity.Vehicle.Trip.ScheduleRelationship == nil {
				trip.ScheduleRelationship = nil
			} else {
				tmp := entity.Vehicle.Trip.GetScheduleRelationship().String()
				trip.ScheduleRelationship = &tmp
			}
		}

		if entity.Vehicle.Vehicle != nil {
			vehicle.Id = entity.Vehicle.Vehicle.Id
			vehicle.Label = entity.Vehicle.Vehicle.Label
			vehicle.LicensePlate = entity.Vehicle.Vehicle.LicensePlate
		}

		if entity.Vehicle.Position != nil {
			position.Latitude = entity.Vehicle.Position.Latitude
			position.Longitude = entity.Vehicle.Position.Longitude
			position.Bearing = entity.Vehicle.Position.Bearing
			position.Odometer = entity.Vehicle.Position.Odometer
			position.Speed = entity.Vehicle.Position.Speed
		}

		currentStopSequence = entity.Vehicle.CurrentStopSequence

		stopId = entity.Vehicle.StopId

		if entity.Vehicle.CurrentStatus == nil {
			currentStatus = nil
		} else {
			tmp := entity.Vehicle.GetCurrentStatus().String()
			currentStatus = &tmp
		}

		timeStamp = entity.Vehicle.Timestamp

		if entity.Vehicle.CongestionLevel == nil {
			congestionLevel = nil
		} else {
			tmp := entity.Vehicle.GetCongestionLevel().String()
			congestionLevel = &tmp
		}

		if entity.Vehicle.OccupancyStatus == nil {
			occupancyStatus = nil
		} else {
			tmp := entity.Vehicle.GetOccupancyStatus().String()
			occupancyStatus = &tmp
		}

		res = append(res, realtime.VehiclePosition{
			Trip:                trip,
			Vehicle:             vehicle,
			Position:            position,
			CurrentStopSequence: currentStopSequence,
			StopId:              stopId,
			CurrentStatus:       currentStatus,
			TimeStamp:           timeStamp,
			CongestionLevel:     congestionLevel,
			OccupancyStatus:     occupancyStatus,
		})
	}
	return res
}

func VehiclePositionPbMap(file string) map[string]realtime.VehiclePosition {
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
	res := make(map[string]realtime.VehiclePosition)

	for _, entity := range feed.Entity {
		var trip realtime.VehiclePositionTripDescriptor
		var vehicle realtime.VehiclePositionVehicleDescriptor
		var position realtime.Position
		var currentStopSequence *uint32
		var stopId *string
		var currentStatus *string
		var timeStamp *uint64
		var congestionLevel *string
		var occupancyStatus *string

		if entity.Vehicle.Trip != nil {
			trip.TripId = entity.Vehicle.Trip.TripId
			trip.RouteId = entity.Vehicle.Trip.RouteId
			trip.DirectionId = entity.Vehicle.Trip.DirectionId
			trip.StartTime = entity.Vehicle.Trip.StartTime
			trip.StartDate = entity.Vehicle.Trip.StartDate
			if entity.Vehicle.Trip.ScheduleRelationship == nil {
				trip.ScheduleRelationship = nil
			} else {
				tmp := entity.Vehicle.Trip.GetScheduleRelationship().String()
				trip.ScheduleRelationship = &tmp
			}
		}

		if entity.Vehicle.Vehicle != nil {
			vehicle.Id = entity.Vehicle.Vehicle.Id
			vehicle.Label = entity.Vehicle.Vehicle.Label
			vehicle.LicensePlate = entity.Vehicle.Vehicle.LicensePlate
		}

		if entity.Vehicle.Position != nil {
			position.Latitude = entity.Vehicle.Position.Latitude
			position.Longitude = entity.Vehicle.Position.Longitude
			position.Bearing = entity.Vehicle.Position.Bearing
			position.Odometer = entity.Vehicle.Position.Odometer
			position.Speed = entity.Vehicle.Position.Speed
		}

		currentStopSequence = entity.Vehicle.CurrentStopSequence

		stopId = entity.Vehicle.StopId

		if entity.Vehicle.CurrentStatus == nil {
			currentStatus = nil
		} else {
			tmp := entity.Vehicle.GetCurrentStatus().String()
			currentStatus = &tmp
		}

		timeStamp = entity.Vehicle.Timestamp

		if entity.Vehicle.CongestionLevel == nil {
			congestionLevel = nil
		} else {
			tmp := entity.Vehicle.GetCongestionLevel().String()
			congestionLevel = &tmp
		}

		if entity.Vehicle.OccupancyStatus == nil {
			occupancyStatus = nil
		} else {
			tmp := entity.Vehicle.GetOccupancyStatus().String()
			occupancyStatus = &tmp
		}

		res[*trip.TripId] = realtime.VehiclePosition{
			Trip:                trip,
			Vehicle:             vehicle,
			Position:            position,
			CurrentStopSequence: currentStopSequence,
			StopId:              stopId,
			CurrentStatus:       currentStatus,
			TimeStamp:           timeStamp,
			CongestionLevel:     congestionLevel,
			OccupancyStatus:     occupancyStatus,
		}
	}
	return res
}
