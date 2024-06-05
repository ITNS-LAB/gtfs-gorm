package ormstatic

import "database/sql"

type Agency struct {
	Id             uint   `gorm:"primaryKey;auto_increment;not null"`
	AgencyId       string `gorm:"index;unique"`
	AgencyName     string `gorm:"not null"`
	AgencyUrl      string `gorm:"not null"`
	AgencyTimezone string `gorm:"not null"`
	AgencyLang     sql.NullString
	AgencyPhone    sql.NullString
	AgencyFareUrl  sql.NullString
	AgencyEmail    sql.NullString
}

func (Agency) TableName() string {
	return "agency"
}

func (a Agency) GetId() any {
	return a.Id
}

func (a Agency) GetAgencyId() any {
	return a.AgencyId
}

func (a Agency) GetAgencyName() any {
	return a.AgencyName
}

func (a Agency) GetAgencyUrl() any {
	return a.AgencyUrl
}

func (a Agency) GetAgencyTimezone() any {
	return a.AgencyTimezone
}

func (a Agency) GetAgencyLang() any {
	if a.AgencyLang.Valid {
		return a.AgencyLang.String
	}
	return nil
}

func (a Agency) GetAgencyPhone() any {
	if a.AgencyPhone.Valid {
		return a.AgencyPhone.String
	}
	return nil
}

func (a Agency) GetAgencyFareUrl() any {
	if a.AgencyFareUrl.Valid {
		return a.AgencyFareUrl.String
	}
	return nil
}

func (a Agency) GetAgencyEmail() any {
	if a.AgencyEmail.Valid {
		return a.AgencyEmail.String
	}
	return nil
}
