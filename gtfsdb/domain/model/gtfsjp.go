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
	Agency        []gtfsjp.AgencyGeom
	AgencyJp      []gtfsjp.AgencyJpGeom
	Stops         []gtfsjp.StopGeom
	Routes        []gtfsjp.RouteGeom
	Trips         []gtfsjp.TripGeom
	OfficeJp      []gtfsjp.OfficeJpGeom
	PatternJp     []gtfsjp.PatternJpGeom
	StopTimes     []gtfsjp.StopTimeGeom
	Calendar      []gtfsjp.CalendarGeom
	CalendarDates []gtfsjp.CalendarDateGeom
	FareAttribute []gtfsjp.FareAttributeGeom
	FareRules     []gtfsjp.FareRuleGeom
	Shapes        []gtfsjp.ShapeGeom
	Frequencies   []gtfsjp.FrequencyGeom
	Transfers     []gtfsjp.TransferGeom
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
