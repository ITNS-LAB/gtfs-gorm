package ormstatic

import "database/sql"

type Translation struct {
	Id int `gorm:"primaryKey;auto_increment;not null"`
	//To avoid duplication with "TableName", "N" should be written in lower case.
	Tablename   string `gorm:"column:table_name;not null"`
	FieldName   string `gorm:"not null"`
	Language    string `gorm:"not null"`
	Translation string `gorm:"not null"`
	RecordId    sql.NullString
	RecordSubId sql.NullString
	FieldValue  sql.NullString
}

func (Translation) TableName() string {
	return "translations"
}

func (t Translation) GetId() any {
	return t.Id
}

func (t Translation) GetTablename() any {
	return t.Tablename
}

func (t Translation) GetFieldName() any {
	return t.FieldName
}

func (t Translation) GetLanguage() any {
	return t.Language
}

func (t Translation) GetTranslation() any {
	return t.Translation
}

func (t Translation) GetRecordId() any {
	if t.RecordId.Valid {
		return t.RecordId.String
	}
	return nil
}

func (t Translation) GetRecordSubId() any {
	if t.RecordSubId.Valid {
		return t.RecordSubId.String
	}
	return nil
}

func (t Translation) GetFieldValue() any {
	if t.FieldValue.Valid {
		return t.FieldValue.String
	}
	return nil
}
