package gtfsschedule

type FareProduct struct {
	FareProductID    string `gorm:"primary_key"`
	FareProductName  *string
	FareMediaID      *string
	Amount           float64            `gorm:"not null"`
	Currency         string             `gorm:"not null"`
	FareLeg          []FareLeg          `gorm:"foreignKey:FareProductID;references:FareProductID "`
	FareTransferRule []FareTransferRule `gorm:"foreignKey:FareProductID;references:FareProductID "`
}
