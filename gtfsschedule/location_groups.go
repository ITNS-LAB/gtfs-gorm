package gtfsschedule

type LocationGroup struct {
	LocationGroupID   string `gorm:"primaryKey"`
	LocationGroupName *string
	StopTimes         []StopTimes         `gorm:"foreignKey:LocationGroupId;references:LocationGroupId "`
	LocationGroupStop []LocationGroupStop `gorm:"foreignKey:LocationGroupId;references:LocationGroupId "`
}
