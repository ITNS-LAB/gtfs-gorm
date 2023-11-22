package static

type FareRule struct {
	Id              int     `gorm:"primaryKey;auto_increment;not null"`
	FareId          *string `gorm:"index;not null"`
	RouteId         *string
	OriginId        *string
	DestinationId   *string
	ContainsId      *string
	OriginStop      Stop `gorm:"foreignKey:OriginId;references:ZoneId"`
	DestinationStop Stop `gorm:"foreignKey:DestinationId;references:ZoneId"`
}

func (FareRule) TableName() string {
	return "fare_rules"
}
