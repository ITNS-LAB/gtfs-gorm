package gtfsjp

import "database/sql"

type AgencyPtr struct {
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

func (AgencyPtr) TableName() string {
	return "agency"
}

type Agency struct {
	Id             uint   `gorm:"primaryKey;auto_increment;not null"`
	AgencyId       string `gorm:"uniqueIndex"`
	AgencyName     string `gorm:"not null"`
	AgencyUrl      string `gorm:"not null"`
	AgencyTimezone string `gorm:"not null"`
	AgencyLang     sql.NullString
	AgencyPhone    sql.NullString
	AgencyFareUrl  sql.NullString
	AgencyEmail    sql.NullString
	Routes         []Route `gorm:"foreignKey:AgencyId;references:AgencyId"`
}

func (Agency) TableName() string {
	return "agency"
}

func (a Agency) GetId() uint {
	return a.Id
}

func (a Agency) GetAgencyId() string {
	return a.AgencyId
}

func (a Agency) GetAgencyName() string {
	return a.AgencyName
}

func (a Agency) GetAgencyUrl() string {
	return a.AgencyUrl
}

func (a Agency) GetAgencyTimezone() string {
	return a.AgencyTimezone
}

func (a Agency) GetAgencyLang(defaultVal string) string {
	if a.AgencyLang.Valid {
		return a.AgencyLang.String
	}
	return defaultVal
}

func (a Agency) GetAgencyPhone(defaultVal string) string {
	if a.AgencyPhone.Valid {
		return a.AgencyPhone.String
	}
	return defaultVal
}

func (a Agency) GetAgencyFareUrl(defaultVal string) string {
	if a.AgencyFareUrl.Valid {
		return a.AgencyFareUrl.String
	}
	return defaultVal
}

func (a Agency) GetAgencyEmail(defaultVal string) string {
	if a.AgencyEmail.Valid {
		return a.AgencyEmail.String
	}
	return defaultVal
}
