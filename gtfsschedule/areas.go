package gtfsschedule

type Areas struct {
	AreaID            string `gorm:"primary_key"`
	AreaName          *string
	FareLegFromAreaID []FareLeg  `gorm:"foreignKey:AreaID;references:FromAreaID "`
	FareLegToAreaID   []FareLeg  `gorm:"foreignKey:AreaID;references:ToAreaID "`
	StopArea          []StopArea `gorm:"foreignKey:AreaID;references:AreaID "`
}
