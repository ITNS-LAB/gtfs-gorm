package ormstatic

import geomdatatypes "github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"

type Stop struct {
	StopId             *string `gorm:"primaryKey"`
	StopCode           *string
	StopName           *string
	StopDesc           *string
	StopLat            *float64
	StopLon            *float64
	ZoneId             *string `gorm:"unique"`
	StopUrl            *string
	LocationType       *int `gorm:"default:0"`
	ParentStation      *string
	StopTimezone       *string
	WheelchairBoarding *int `gorm:"default:0"`
	LevelId            *string
	PlatformCode       *string
	Geom               *geomdatatypes.Geometry `gorm:"index"`
	StopTimes          []StopTime              `gorm:"foreignKey:StopId;references:StopId"`
}

func (Stop) TableName() string {
	return "stops"
}
