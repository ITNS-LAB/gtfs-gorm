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
}

type GtfsJpGeom struct {
	Agency        []gtfsjp.Agency
	AgencyJp      []gtfsjp.AgencyJp
	StopsGeom     []gtfsjp.StopGeom
	Routes        []gtfsjp.Route
	TripsGeom     []gtfsjp.TripGeom
	OfficeJp      []gtfsjp.OfficeJp
	PatternJp     []gtfsjp.PatternJp
	StopTimes     []gtfsjp.StopTime
	Calendar      []gtfsjp.Calendar
	CalendarDates []gtfsjp.CalendarDate
	FareAttribute []gtfsjp.FareAttribute
	FareRules     []gtfsjp.FareRule
	ShapesGeom    []gtfsjp.ShapeGeom
	Frequencies   []gtfsjp.Frequency
	Transfers     []gtfsjp.Transfer
	FeedInfo      []gtfsjp.FeedInfo
	Translations  []gtfsjp.Translation
}

type Trip struct {
	gtfsjp.Trip
}

type TripGeom struct {
	gtfsjp.TripGeom
}

type Shape struct {
	gtfsjp.Shape
}

type ShapeGeom struct {
	gtfsjp.ShapeGeom
}

type ShapeEx struct {
	gtfsjp.ShapeEx
}

type ShapeDetail struct {
	gtfsjp.ShapeDetail
}

type StopTime struct {
	gtfsjp.StopTime
}
