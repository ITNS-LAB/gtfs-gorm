package model

import "github.com/ITNS-LAB/gtfs-gorm/ormstatic"

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
}
