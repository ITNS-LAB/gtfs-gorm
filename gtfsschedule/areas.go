package gtfsschedule

type Areas struct {
	AreaID   string `gorm:"primary_key"`
	AreaName *string
}
