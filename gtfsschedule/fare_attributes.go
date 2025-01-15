package gtfsschedule

type FareAttributes struct {
	FareID        int    `gorm:"primary_key"`
	Price         int    `gorm:"not null"`
	CurrencyType  string `gorm:"not null"`
	PaymentMethod int    `gorm:"not null"`
	Transfers     *int   `gorm:"not null"`
	AgencyID      int
}
