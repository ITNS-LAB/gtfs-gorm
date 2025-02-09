package model

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsschedule"
)

type GtfsShedule struct {
	Agency            []gtfsschedule.Agency
	Areas             []gtfsschedule.Areas
	Attributions      []gtfsschedule.Attribution
	BookingRules      []gtfsschedule.BookingRule
	Calendar          []gtfsschedule.Calendar
	CalendarDates     []gtfsschedule.CalendarDates
	FareAttributes    []gtfsschedule.FareAttributes
	FareLegJoinRules  []gtfsschedule.FareLegJoinRules
	FareLegRules      []gtfsschedule.FareLegRules
	FareMedia         []gtfsschedule.FareMedia
	FareProduct       []gtfsschedule.FareProduct
	FareRules         []gtfsschedule.FareRules
	FareTransferRule  []gtfsschedule.FareTransferRule
	FeedInfo          []gtfsschedule.FeedInfo
	Frequencies       []gtfsschedule.Frequencies
	Levels            []gtfsschedule.Levels
	LocationGroupStop []gtfsschedule.LocationGroupStop
	LocationGroup     []gtfsschedule.LocationGroup
	Network           []gtfsschedule.Network
	Pathway           []gtfsschedule.Pathway
	RouteNetwork      []gtfsschedule.RouteNetwork
	Route             []gtfsschedule.Route
	Shape             []gtfsschedule.Shape
	StopArea          []gtfsschedule.StopArea
	StopTimes         []gtfsschedule.StopTimes
	Stop              []gtfsschedule.Stop
	TimeFrame         []gtfsschedule.TimeFrame
	Transfer          []gtfsschedule.Transfer
	Translation       []gtfsschedule.Translation
	Trips             []gtfsschedule.Trips
}

type Trip struct {
	gtfsschedule.Trips
}

/*
type TripGeom struct {
	gtfsjp.TripGeom
}

*/

type Shape struct {
	gtfsschedule.Shape
}

/*
	type ShapeGeom struct {
		gtfsjp.ShapeGeom
	}
*/
type ShapeEx struct {
	gtfsschedule.ShapeEx
}

/*
type ShapeExGeom struct {
	gtfsjp.ShapeExGeom
}

*/

type ShapeDetail struct {
	gtfsschedule.ShapeDetail
}

/*
type ShapeDetailGeom struct {
	gtfsjp.ShapeDetailGeom
}

*/

type StopTime struct {
	gtfsschedule.StopTimes
}

type TripWithStopLocation struct {
	TripId       string
	StopId       string
	StopSequence int
	StopLat      float64
	StopLon      float64
}
