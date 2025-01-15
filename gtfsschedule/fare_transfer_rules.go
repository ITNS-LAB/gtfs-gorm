package gtfsschedule

type FareTransferRule struct {
	FromLegGroupID    *string
	ToLegGroupID      *string
	TransferCount     *int
	DurationLimit     *int
	DurationLimitType *int
	FareTransferType  int `gorm:"not null"`
	FareProductID     *string
}
