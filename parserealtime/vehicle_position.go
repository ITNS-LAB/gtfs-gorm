package parserealtime

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormrealtime"
	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
	"os"
)

func VehiclePositionPbSlice(file string) ([]ormrealtime.VehiclePosition, error) {
	var res []ormrealtime.VehiclePosition

	// gtfs-ormrealtime ファイルをバイナリ形式で読み込む
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
		var trip ormrealtime.VehiclePositionTripDescriptor
		var vehicle ormrealtime.VehiclePositionVehicleDescriptor
		var position ormrealtime.Position
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

		res = append(res, ormrealtime.VehiclePosition{
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
	return res, nil
}

func VehiclePositionPbMap(file string) (map[string]ormrealtime.VehiclePosition, error) {
	// mapの初期化
	res := make(map[string]ormrealtime.VehiclePosition)

	// gtfs-ormrealtime ファイルをバイナリ形式で読み込む
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
		var trip ormrealtime.VehiclePositionTripDescriptor
		var vehicle ormrealtime.VehiclePositionVehicleDescriptor
		var position ormrealtime.Position
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

		res[*trip.TripId] = ormrealtime.VehiclePosition{
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
	return res, nil
}
