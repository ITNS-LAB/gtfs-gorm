package gtfsschedule

type FareProduct struct {
	FareProductID   string `gorm:"primary_key"`
	FareProductName *string
	FareMediaID     *string
	Amount          float64 `gorm:"not null"`
	Currency        string  `gorm:"not null"`
}
