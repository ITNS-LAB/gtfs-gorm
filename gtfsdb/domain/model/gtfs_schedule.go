package model

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
)

type GtfsSchedule struct {
	Agency        []gtfsjp.Agency
	Stops         []gtfsjp.Stop
	Routes        []gtfsjp.Route
	Trips         []gtfsjp.Trip
	StopTimes     []gtfsjp.StopTime
	Calendar      []gtfsjp.Calendar
	CalendarDates []gtfsjp.CalendarDate
	FareAttribute []gtfsjp.FareAttribute
	FareRules     []gtfsjp.FareRule
	// FareMedia
	// FareProducts
	// FareLegRules
	// FareTransferRules
	// Areas
	// StopAreas
	Shapes       []gtfsjp.Shape
	Frequencies  []gtfsjp.Frequency
	Transfers    []gtfsjp.Transfer
	Pathways     []gtfsjp.Pathway
	Levels       []gtfsjp.Level
	Translations []gtfsjp.Translation
	FeedInfo     []gtfsjp.FeedInfo
	Attributions []gtfsjp.Attribution
	ShapesEx     []gtfsjp.ShapeEx
	ShapesDetail []gtfsjp.ShapeDetail
}

type StopTimeWithLocation struct {
	TripId       *string `gorm:"primaryKey"`
	StopId       *string `gorm:"primaryKey"`
	StopSequence *int    `gorm:"primaryKey"`
	StopLat      *float64
	StopLon      *float64
}
