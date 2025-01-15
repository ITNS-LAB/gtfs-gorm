package gtfsschedule

type Transfer struct {
	FromStopID      string
	ToStopID        string
	FromRouteID     *string
	ToRouteID       *string
	FromTripID      *string
	ToTripID        *string
	TransferType    int `gorm:"not null"` // 接続タイプを示します (0, 1, 2, 3, 4, 5)
	MinTransferTime *int
}
