package ormstatic

type Agency struct {
	Id             int     `gorm:"primaryKey;auto_increment"`
	AgencyId       *string `gorm:"index;unique"`
	AgencyName     *string `gorm:"not null"`
	AgencyUrl      *string `gorm:"not null"`
	AgencyTimezone *string `gorm:"not null"`
	AgencyLang     *string
	AgencyPhone    *string
	AgencyFareUrl  *string
	AgencyEmail    *string
	Routes         []Route `gorm:"foreignKey:AgencyId;references:AgencyId"`
}

func (Agency) TableName() string {
	return "agency"
}
