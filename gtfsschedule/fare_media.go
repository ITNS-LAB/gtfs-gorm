package gtfsschedule

type FareMedia struct {
	FareMediaID   string `gorm:"primary_key"`
	FareMediaName *string
	FareMediaType int `gorm:"not null"`
}
