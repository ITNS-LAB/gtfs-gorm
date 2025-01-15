package gtfsschedule

import "gorm.io/datatypes"

type TimeFrame struct {
	TimeframeGroupId            int `gorm:"primary_key"`
	StartTime                   *datatypes.Time
	EndTime                     *datatypes.Time
	ServiceId                   int       `gorm:"not null"`
	FareLegFromTimeframeGroupID []FareLeg `gorm:"foreignKey:TimeframeGroupId;references:FromTimeframeGroupID "`
	FareLegToTimeframeGroupID   []FareLeg `gorm:"foreignKey:TimeframeGroupId;references:ToTimeframeGroupID "`
}
