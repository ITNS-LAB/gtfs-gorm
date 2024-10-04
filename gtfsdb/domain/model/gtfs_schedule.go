package model

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
)

type GtfsSchedule struct {
	Agency        []ormstatic.Agency
	Stops         []ormstatic.Stop
	Routes        []ormstatic.Route
	Trips         []ormstatic.Trip
	StopTimes     []ormstatic.StopTime
	Calendar      []ormstatic.Calendar
	CalendarDates []ormstatic.CalendarDate
	FareAttribute []ormstatic.FareAttribute
	FareRules     []ormstatic.FareRule
	// FareMedia
	// FareProducts
	// FareLegRules
	// FareTransferRules
	// Areas
	// StopAreas
	Shapes       []ormstatic.Shape
	Frequencies  []ormstatic.Frequency
	Transfers    []ormstatic.Transfer
	Pathways     []ormstatic.Pathway
	Levels       []ormstatic.Level
	Translations []ormstatic.Translation
	FeedInfo     []ormstatic.FeedInfo
	Attributions []ormstatic.Attribution
	ShapesEx     []ormstatic.ShapeEx
	ShapesDetail []ormstatic.ShapeDetail
}

type StopTimeWithLocation struct {
	TripId       *string `gorm:"primaryKey"`
	StopId       *string `gorm:"primaryKey"`
	StopSequence *int    `gorm:"primaryKey"`
	StopLat      *float64
	StopLon      *float64
}
