package gtfsjp

type Transfer struct {
	Id              int     `gorm:"primaryKey;auto_increment"`
	FromStopId      *string `gorm:"primaryKey"`
	ToStopId        *string `gorm:"primaryKey"`
	FromRouteId     *string
	ToRouteId       *string
	FromTripId      *string
	ToTripId        *string
	TransferType    *int `gorm:"not null"`
	MinTransferTime *int
	FromStop        Stop `gorm:"foreignKey:FromStopId;references:StopId"`
	ToStop          Stop `gorm:"foreignKey:ToStopId;references:StopId"`
	FromTrip        Trip `gorm:"foreignKey:FromTripId;references:TripId"`
	ToTrip          Trip `gorm:"foreignKey:ToTripId;references:TripId"`
}

func (Transfer) TableName() string {
	return "transfers"
}
