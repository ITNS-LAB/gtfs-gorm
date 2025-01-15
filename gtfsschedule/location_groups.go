package gtfsschedule

type LocationGroup struct {
	LocationGroupID   string `gorm:"primaryKey"`
	LocationGroupName *string
}
