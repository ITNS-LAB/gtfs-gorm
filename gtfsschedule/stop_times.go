package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type StopTimes struct {
	TripId                   string         `gorm:"primaryKey"`
	ArrivalTime              datatypes.Time `gorm:"index;not null"`
	DepartureTime            datatypes.Time `gorm:"index;not null"`
	StopId                   string
	LocationGroupId          *string
	LocationId               *string
	StopSequence             int `gorm:"primaryKey"`
	StopHeadsign             *string
	StartPickupDropOffWindow *datatypes.Time
	EndPickupDropOffWindow   *datatypes.Time
	PickupType               *int
	DropOffType              *int
	ContinuousPickup         *int
	ContinuousDropOff        *int
	ShapeDistTraveled        *float64
	Timepoint                *int
	PickupBookingRuleId      *string
	DropOffBookingRuleid     *string
}

func ParseStopTimes(path string) ([]StopTimes, error) {
	// Open the CSV
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open StopTimes CSV: %w", err)
	}

	// Parse the data and create a slice of StopTimes structs
	var stopTimesList []StopTimes
	for i := 0; i < len(df.Records); i++ {
		tripId, err := df.GetString(i, "trip_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'trip_id' at row %d: %w", i, err)
		}

		arrivalTime, err := df.GetTime(i, "arrival_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'arrival_time' at row %d: %w", i, err)
		}

		departureTime, err := df.GetTime(i, "departure_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'departure_time' at row %d: %w", i, err)
		}

		stopId, err := df.GetString(i, "stop_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_id' at row %d: %w", i, err)
		}

		locationGroupId, err := df.GetStringPtr(i, "location_group_id")
		if err != nil {
			locationGroupId = nil // Set to nil if no data
		}

		locationId, err := df.GetStringPtr(i, "location_id")
		if err != nil {
			locationId = nil // Set to nil if no data
		}

		stopSequence, err := df.GetInt(i, "stop_sequence")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'stop_sequence' at row %d: %w", i, err)
		}

		stopHeadsign, err := df.GetStringPtr(i, "stop_headsign")
		if err != nil {
			stopHeadsign = nil // Set to nil if no data
		}

		startPickupDropOffWindow, err := df.GetTimePtr(i, "start_pickup_dropoff_window")
		if err != nil {
			startPickupDropOffWindow = nil // Set to nil if no data
		}

		endPickupDropOffWindow, err := df.GetTimePtr(i, "end_pickup_dropoff_window")
		if err != nil {
			endPickupDropOffWindow = nil // Set to nil if no data
		}

		pickupType, err := df.GetIntPtr(i, "pickup_type")
		if err != nil {
			pickupType = nil // Set to nil if no data
		}

		dropOffType, err := df.GetIntPtr(i, "drop_off_type")
		if err != nil {
			dropOffType = nil // Set to nil if no data
		}

		continuousPickup, err := df.GetIntPtr(i, "continuous_pickup")
		if err != nil {
			continuousPickup = nil // Set to nil if no data
		}

		continuousDropOff, err := df.GetIntPtr(i, "continuous_drop_off")
		if err != nil {
			continuousDropOff = nil // Set to nil if no data
		}

		shapeDistTraveled, err := df.GetFloatPtr(i, "shape_dist_traveled")
		if err != nil {
			shapeDistTraveled = nil // Set to nil if no data
		}

		timepoint, err := df.GetIntPtr(i, "timepoint")
		if err != nil {
			timepoint = nil // Set to nil if no data
		}

		pickupBookingRuleId, err := df.GetStringPtr(i, "pickup_booking_rule_id")
		if err != nil {
			pickupBookingRuleId = nil // Set to nil if no data
		}

		dropOffBookingRuleid, err := df.GetStringPtr(i, "drop_off_booking_rule_id")
		if err != nil {
			dropOffBookingRuleid = nil // Set to nil if no data
		}

		// Add the StopTimes struct to the list
		stopTimesList = append(stopTimesList, StopTimes{
			TripId:                   tripId,
			ArrivalTime:              arrivalTime,
			DepartureTime:            departureTime,
			StopId:                   stopId,
			LocationGroupId:          locationGroupId,
			LocationId:               locationId,
			StopSequence:             stopSequence,
			StopHeadsign:             stopHeadsign,
			StartPickupDropOffWindow: startPickupDropOffWindow,
			EndPickupDropOffWindow:   endPickupDropOffWindow,
			PickupType:               pickupType,
			DropOffType:              dropOffType,
			ContinuousPickup:         continuousPickup,
			ContinuousDropOff:        continuousDropOff,
			ShapeDistTraveled:        shapeDistTraveled,
			Timepoint:                timepoint,
			PickupBookingRuleId:      pickupBookingRuleId,
			DropOffBookingRuleid:     dropOffBookingRuleid,
		})
	}

	return stopTimesList, nil
}
