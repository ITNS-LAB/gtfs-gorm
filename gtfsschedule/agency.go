package gtfsschedule

type Agency struct {
	agencyID       string `gorm:"primary_key"`
	agencyName     string `gorm:"not null"`
	agencyUrl      string `gorm:"not null"`
	agencyTimezone string `gorm:"not null"`
	agencyLang     *string
	agencyPhone    *string
	agencyFareUrl  *string
	agencyEmail    *string
}
