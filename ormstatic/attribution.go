package ormstatic

type Attribution struct {
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

func (Attribution) TableName() string {
	return "attributions"
}
