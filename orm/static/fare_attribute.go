package static

type FareAttribute struct {
	FareId           *string  `gorm:"primaryKey;index;not null"`
	Price            *float64 `gorm:"index;not null"`
	CurrencyType     *string  `gorm:"not null"`
	PaymentMethod    *int     `gorm:"not null"`
	Transfers        *int     `gorm:"not null"`
	AgencyId         *string
	TransferDuration *int
	FareRule         FareRule `gorm:"foreignKey:FareId"`
}

func (FareAttribute) TableName() string {
	return "fare_attributes"
}
