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

type GtfsSheduleGeom struct {
	Agency            []gtfsschedule.AgencyGeom
	Areas             []gtfsschedule.AreasGeom
	Attributions      []gtfsschedule.AttributionGeom
	BookingRules      []gtfsschedule.BookingRuleGeom
	Calendar          []gtfsschedule.CalendarGeom
	CalendarDates     []gtfsschedule.CalendarDatesGeom
	FareAttributes    []gtfsschedule.FareAttributesGeom
	FareLegJoinRules  []gtfsschedule.FareLegJoinRulesGeom
	FareLegRules      []gtfsschedule.FareLegRulesGeom
	FareMedia         []gtfsschedule.FareMediaGeom
	FareProduct       []gtfsschedule.FareProductGeom
	FareRules         []gtfsschedule.FareRulesGeom
	FareTransferRule  []gtfsschedule.FareTransferRuleGeom
	FeedInfo          []gtfsschedule.FeedInfoGeom
	Frequencies       []gtfsschedule.FrequenciesGeom
	Levels            []gtfsschedule.LevelsGeom
	LocationGroupStop []gtfsschedule.LocationGroupStopGeom
	LocationGroup     []gtfsschedule.LocationGroupGeom
	Network           []gtfsschedule.NetworkGeom
	Pathway           []gtfsschedule.PathwayGeom
	RouteNetwork      []gtfsschedule.RouteNetworkGeom
	Route             []gtfsschedule.RouteGeom
	Shape             []gtfsschedule.ShapeGeom
	StopArea          []gtfsschedule.StopAreaGeom
	StopTimes         []gtfsschedule.StopTimesGeom
	Stop              []gtfsschedule.StopGeom
	TimeFrame         []gtfsschedule.TimeFrameGeom
	Transfer          []gtfsschedule.TransferGeom
	Translation       []gtfsschedule.TranslationGeom
	Trips             []gtfsschedule.TripsGeom
}

type Trip struct {
	gtfsschedule.Trips
}

type TripGeom struct {
	gtfsschedule.TripsGeom
}

type Shape struct {
	gtfsschedule.Shape
}

type ShapeGeom struct {
	gtfsschedule.ShapeGeom
}

type ShapeEx struct {
	gtfsschedule.ShapeEx
}

type ShapeExGeom struct {
	gtfsschedule.ShapeExGeom
}

type ShapeDetail struct {
	gtfsschedule.ShapeDetail
}

type ShapeDetailGeom struct {
	gtfsschedule.ShapeDetailGeom
}

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
