package static

type Attribution struct {
	AttributionId    *string `gorm:"primaryKey;not null"`
	AgencyId         *string `gorm:"primaryKey;not null"`
	RouteId          *string `gorm:"primaryKey;not null"`
	TripId           *string `gorm:"primaryKey;not null"`
	OrganizationName *string
	IsProducer       *int `gorm:"default:0"`
	IsOperator       *int `gorm:"default:0"`
	IsAuthority      *int `gorm:"default:0"`
	AttributionUrl   *string
	AttributionEmail *string
	AttributionPhone *string
}

func (Attribution) TableName() string {
	return "attributions"
}
