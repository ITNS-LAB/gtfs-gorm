package gtfsjp

type Translation struct {
	Id int `gorm:"primaryKey;auto_increment"`
	//To avoid duplication with "TableName", "N" should be written in lower case.
	Tablename   *string `gorm:"column:table_name;not null"`
	FieldName   *string `gorm:"not null"`
	Language    *string `gorm:"not null"`
	Translation *string `gorm:"not null"`
	RecordId    *string
	RecordSubId *string
	FieldValue  *string
}

func (Translation) TableName() string {
	return "translations"
}
