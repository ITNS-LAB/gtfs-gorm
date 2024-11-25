package gtfsjp

import "database/sql"

type AttributionPtr struct {
	AttributionId    *string `gorm:"primaryKey"`
	AgencyId         *string `gorm:"primaryKey"`
	RouteId          *string `gorm:"primaryKey"`
	TripId           *string `gorm:"primaryKey"`
	OrganizationName *string `gorm:"not null"`
	IsProducer       *int    `gorm:"default:0"`
	IsOperator       *int    `gorm:"default:0"`
	IsAuthority      *int    `gorm:"default:0"`
	AttributionUrl   *string
	AttributionEmail *string
	AttributionPhone *string
}

func (AttributionPtr) TableName() string {
	return "attributions"
}

type Attribution struct {
	AttributionId    string        `gorm:"primaryKey"`
	AgencyId         string        `gorm:"primaryKey"`
	RouteId          string        `gorm:"primaryKey"`
	TripId           string        `gorm:"primaryKey"`
	OrganizationName string        `gorm:"not null"`
	IsProducer       sql.NullInt64 `gorm:"default:0"`
	IsOperator       sql.NullInt64 `gorm:"default:0"`
	IsAuthority      sql.NullInt64 `gorm:"default:0"`
	AttributionUrl   sql.NullString
	AttributionEmail sql.NullString
	AttributionPhone sql.NullString
}

func (Attribution) TableName() string {
	return "attributions"
}

func (a Attribution) GetAttributionId() string {
	return a.AttributionId
}

func (a Attribution) GetAgencyId() string {
	return a.AgencyId
}

func (a Attribution) GetRouteId() string {
	return a.RouteId
}

func (a Attribution) GetTripId() string {
	return a.TripId
}

func (a Attribution) GetOrganizationName() string {
	return a.OrganizationName
}

func (a Attribution) GetIsProducer(defaultVal int64) int64 {
	if a.IsProducer.Valid {
		return a.IsProducer.Int64
	}
	return defaultVal
}

func (a Attribution) GetIsOperator(defaultVal int64) int64 {
	if a.IsOperator.Valid {
		return a.IsOperator.Int64
	}
	return defaultVal
}

func (a Attribution) GetIsAuthority(defaultVal int64) int64 {
	if a.IsAuthority.Valid {
		return a.IsAuthority.Int64
	}
	return defaultVal
}

func (a Attribution) GetAttributionUrl(defaultVal string) string {
	if a.AttributionUrl.Valid {
		return a.AttributionUrl.String
	}
	return defaultVal
}

func (a Attribution) GetAttributionEmail(defaultVal string) string {
	if a.AttributionEmail.Valid {
		return a.AttributionEmail.String
	}
	return defaultVal
}

func (a Attribution) GetAttributionPhone(defaultVal string) string {
	if a.AttributionPhone.Valid {
		return a.AttributionPhone.String
	}
	return defaultVal
}
