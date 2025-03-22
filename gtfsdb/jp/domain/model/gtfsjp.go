package model

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"
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

type TripGeomLine struct {
	TripId string
	Geom   gormdatatypes.Geometry
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

type ShapeExGeom struct {
	gtfsjp.ShapeExGeom
}

type ShapeDetail struct {
	gtfsjp.ShapeDetail
}

type ShapeDetailGeom struct {
	gtfsjp.ShapeDetailGeom
}

type StopTime struct {
	gtfsjp.StopTime
}

type TripWithStopLocation struct {
	TripId       string
	StopId       string
	StopSequence int
	StopLat      float64
	StopLon      float64
}
