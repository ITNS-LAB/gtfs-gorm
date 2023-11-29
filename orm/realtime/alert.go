package realtime

type Alert struct {
	AlertId        uint             `gorm:"primaryKey;auto_increment"`
	ActivePeriod   []TimeRange      `gorm:"foreignKey:AlertId"`
	InformedEntity []EntitySelector `gorm:"foreignKey:AlertId"`
	Cause          *string
	Effect         *string
	Url            []UrlTranslation         `gorm:"foreignKey:AlertId"`
	HeaderText     []HeaderTextTranslation  `gorm:"foreignKey:AlertId"`
	Description    []DescriptionTranslation `gorm:"foreignKey:AlertId"`
}

func (Alert) TableName() string {
	return "alert.alert"
}

type TimeRange struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Start    *uint64
	End      *uint64
}

func (TimeRange) TableName() string {
	return "alert.active_period"
}

type EntitySelector struct {
	InformedEntityId uint `gorm:"primaryKey;auto_increment"`
	AlertId          uint
	AgencyId         *string
	RouteId          *string
	RouteType        *int32
	DirectionId      *uint32
	Trip             AlertTripDescriptor `gorm:"foreignKey:InformedEntityId"`
	StopId           *string
}

func (EntitySelector) TableName() string {
	return "alert.informed_entity"
}

type AlertTripDescriptor struct {
	SerialId             uint `gorm:"primaryKey;auto_increment"`
	InformedEntityId     uint
	TripId               *string
	RouteId              *string
	DirectionId          *uint32
	StartTime            *string
	StartDate            *string
	ScheduleRelationship *string
}

func (AlertTripDescriptor) TableName() string {
	return "alert.trip"
}

type UrlTranslation struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Text     *string
	Language *string
}

func (UrlTranslation) TableName() string {
	return "alert.url"
}

type HeaderTextTranslation struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Text     *string
	Language *string
}

func (HeaderTextTranslation) TableName() string {
	return "alert.header_text"
}

type DescriptionTranslation struct {
	SerialId uint `gorm:"primaryKey;auto_increment"`
	AlertId  uint
	Text     *string
	Language *string
}

func (DescriptionTranslation) TableName() string {
	return "alert.description_text"
}
