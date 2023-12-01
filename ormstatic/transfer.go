package ormstatic

type Transfer struct {
	Id              int     `gorm:"primaryKey;auto_increment;not null"`
	FromStopId      *string `gorm:"primaryKey;not null"`
	ToStopId        *string `gorm:"primaryKey;not null"`
	TransferType    *int    `gorm:"not null"`
	MinTransferTime *int
	FromStop        Stop `gorm:"foreignKey:FromStopId;references:StopId"`
	ToStop          Stop `gorm:"foreignKey:ToStopId;references:StopId"`
}

func (Transfer) TableName() string {
	return "transfers"
}
