package gtfsschedule

type Stop struct {
	StopId                     string `gorm:"primary_key"`
	StopCode                   *string
	StopName                   string `gorm:"not null"`
	TtsStopName                *string
	StopDesc                   *string
	StopLat                    float64 `gorm:"not null"`
	StopLon                    float64 `gorm:"not null"`
	ZoneId                     *string
	StopUrl                    *string
	LocationType               *int
	ParentStation              *string `gorm:"not null"`
	StopTimezone               *string
	WheelchairBoarding         *int
	LevelId                    *string
	PlatformCode               *string
	StopTimes                  []StopTimes         `gorm:"foreignKey:StopId;references:StopId"`
	FareRulesOriginID          FareRules           `gorm:"foreignKey:ZoneId;references:OriginID"`
	FareRulesDestinationID     FareRules           `gorm:"foreignKey:ZoneId;references:DestinationID"`
	FareRulesContainsId        FareRules           `gorm:"foreignKey:ZoneId;references:ContainsId"`
	FareLegJoinRulesFromStopID FareLegJoinRules    `gorm:"foreignKey:StopId;references:FromStopID "`
	FareLegJoinRulesToStopID   FareLegJoinRules    `gorm:"foreignKey:StopId;references:ToStopID "`
	StopArea                   []StopArea          `gorm:"foreignKey:StopId;references:StopId "`
	TransferFromStopID         []Transfer          `gorm:"foreignKey:StopId;references:FromStopID "`
	TransferToStopID           []Transfer          `gorm:"foreignKey:StopId;references:ToStopID "`
	PathwayFromStopID          []Pathway           `gorm:"foreignKey:StopId;references:FromStopID "`
	PathwayToStopID            []Pathway           `gorm:"foreignKey:StopId;references:ToStopID "`
	LocationGroupStop          []LocationGroupStop `gorm:"foreignKey:StopId;references:StopId "`
}
