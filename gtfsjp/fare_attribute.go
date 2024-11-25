package gtfsjp

type FareAttribute struct {
	FareId           *string  `gorm:"primaryKey;index"`
	Price            *float64 `gorm:"index;not null"`
	CurrencyType     *string  `gorm:"primaryKey"`
	PaymentMethod    *int     `gorm:"not null"`
	Transfers        *int     //必須の項目だが、空の場合があるため、nullを許容
	AgencyId         *string
	TransferDuration *int
	//FareRule         FareRule `gorm:"foreignKey:FareId"`
}

func (FareAttribute) TableName() string {
	return "fare_attributes"
}
