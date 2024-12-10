package model

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
)

type GtfsJp struct {
	Agency        []gtfsjp.Agency
	AgencyJp      []gtfsjp.AgencyJp
	Stops         []gtfsjp.Stop
	Routes        []gtfsjp.Route
	Trips         []gtfsjp.Trip
	OfficeJp      []gtfsjp.OfficeJp
	PatternJp     []gtfsjp.PatternJp
	StopTimes     []gtfsjp.StopTime
	Calendar      []gtfsjp.Calendar
	CalendarDates []gtfsjp.CalendarDate
	FareAttribute []gtfsjp.FareAttribute
	FareRules     []gtfsjp.FareRule
	Shapes        []gtfsjp.Shape
	Frequencies   []gtfsjp.Frequency
	Transfers     []gtfsjp.Transfer
	FeedInfo      []gtfsjp.FeedInfo
	Translations  []gtfsjp.Translation
	ShapesEx      []gtfsjp.ShapeEx
	ShapesDetail  []gtfsjp.ShapeDetail
}

type StopTimeWithLocation struct {
	TripId       *string `gorm:"primaryKey"`
	StopId       *string `gorm:"primaryKey"`
	StopSequence *int    `gorm:"primaryKey"`
	StopLat      *float64
	StopLon      *float64
}
