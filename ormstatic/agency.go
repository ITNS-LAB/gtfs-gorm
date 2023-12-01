package ormstatic

type Agency struct {
	Id             int     `gorm:"primaryKey;auto_increment;not null"`
	AgencyId       *string `gorm:"index;unique"`
	AgencyName     *string `gorm:"not null"`
	AgencyUrl      *string `gorm:"not null"`
	AgencyTimezone *string `gorm:"not null"`
	AgencyLang     *string
	AgencyPhone    *string
	AgencyFareUrl  *string
	AgencyEmail    *string
}

func (Agency) TableName() string {
	return "agency"
}
