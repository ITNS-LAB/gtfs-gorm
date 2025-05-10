package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
	"gorm.io/datatypes"
)

type TimeFrame struct {
	TimeframeGroupId int `gorm:"primaryKey"`
	StartTime        *datatypes.Time
	EndTime          *datatypes.Time
	ServiceId        string `gorm:"not null"`
	//FareLegFromTimeframeGroupID []FareLeg `gorm:"foreignKey:TimeframeGroupId;references:FromTimeframeGroupID"`
	//FareLegToTimeframeGroupID   []FareLeg `gorm:"foreignKey:TimeframeGroupId;references:ToTimeframeGroupID"`
}

func (TimeFrame) TableName() string {
	return "timeframe"
}

func ParseTimeFrame(path string) ([]TimeFrame, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open timeframe CSV: %w", err)
	}

	// Parse the data and create a slice of TimeFrame structs
	var timeFrames []TimeFrame
	for i := 0; i < len(df.Records); i++ {
		// Read fields for TimeFrame
		timeframeGroupId, err := df.GetInt(i, "timeframe_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'timeframe_group_id' at row %d: %w", i, err)
		}

		startTime, err := df.GetTimePtr(i, "start_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'start_time' at row %d: %w", i, err)
		}

		endTime, err := df.GetTimePtr(i, "end_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'end_time' at row %d: %w", i, err)
		}

		serviceId, err := df.GetString(i, "service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'service_id' at row %d: %w", i, err)
		}

		// Create the TimeFrame struct and append to the list
		timeFrames = append(timeFrames, TimeFrame{
			TimeframeGroupId: timeframeGroupId,
			StartTime:        startTime,
			EndTime:          endTime,
			ServiceId:        serviceId,
		})
	}

	return timeFrames, nil
}

type TimeFrameGeom struct {
	TimeframeGroupId int `gorm:"primaryKey"`
	StartTime        *datatypes.Time
	EndTime          *datatypes.Time
	ServiceId        string `gorm:"not null"`
	//FareLegFromTimeframeGroupID []FareLegGeom `gorm:"foreignKey:TimeframeGroupId;references:FromTimeframeGroupID"`
	//FareLegToTimeframeGroupID   []FareLegGeom `gorm:"foreignKey:TimeframeGroupId;references:ToTimeframeGroupID"`
}

func (TimeFrameGeom) TableName() string {
	return "timeframe"
}

func ParseTimeFrameGeom(path string) ([]TimeFrameGeom, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open timeframe CSV: %w", err)
	}

	// Parse the data and create a slice of TimeFrame structs
	var timeFrames []TimeFrameGeom
	for i := 0; i < len(df.Records); i++ {
		// Read fields for TimeFrame
		timeframeGroupId, err := df.GetInt(i, "timeframe_group_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'timeframe_group_id' at row %d: %w", i, err)
		}

		startTime, err := df.GetTimePtr(i, "start_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'start_time' at row %d: %w", i, err)
		}

		endTime, err := df.GetTimePtr(i, "end_time")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'end_time' at row %d: %w", i, err)
		}

		serviceId, err := df.GetString(i, "service_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'service_id' at row %d: %w", i, err)
		}

		// Create the TimeFrame struct and append to the list
		timeFrames = append(timeFrames, TimeFrameGeom{
			TimeframeGroupId: timeframeGroupId,
			StartTime:        startTime,
			EndTime:          endTime,
			ServiceId:        serviceId,
		})
	}

	return timeFrames, nil
}
