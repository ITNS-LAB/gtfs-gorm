package gtfsjp

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/csvutil"
)

type Translation struct {
	Id int `gorm:"primaryKey;auto_increment"`
	//To avoid duplication with "TableName", "N" should be written in lower case.
	Tablename   string `gorm:"column:table_name;not null"`
	FieldName   string `gorm:"not null"`
	Language    string `gorm:"not null"`
	Translation string `gorm:"not null"`
	RecordId    *string
	RecordSubId *string
	FieldValue  *string
}

func (Translation) TableName() string {
	return "translations"
}

func ParseTranslations(path string) ([]Translation, error) {
	// CSV を開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open translations CSV: %w", err)
	}

	// データを解析して Translation 構造体のスライスを作成
	var translations []Translation
	for i := 0; i < len(df.Records); i++ {
		tablename, err := df.GetString(i, "table_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'table_name' at row %d: %w", i, err)
		}

		fieldName, err := df.GetString(i, "field_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'field_name' at row %d: %w", i, err)
		}

		language, err := df.GetString(i, "language")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'language' at row %d: %w", i, err)
		}

		translation, err := df.GetString(i, "translation")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'translation' at row %d: %w", i, err)
		}

		recordId, err := df.GetStringPtr(i, "record_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'record_id' at row %d: %w", i, err)
		}

		recordSubId, err := df.GetStringPtr(i, "record_sub_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'record_sub_id' at row %d: %w", i, err)
		}

		fieldValue, err := df.GetStringPtr(i, "field_value")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'field_value' at row %d: %w", i, err)
		}

		// Translation 構造体を作成しリストに追加
		translations = append(translations, Translation{
			Tablename:   tablename,
			FieldName:   fieldName,
			Language:    language,
			Translation: translation,
			RecordId:    recordId,
			RecordSubId: recordSubId,
			FieldValue:  fieldValue,
		})
	}

	return translations, nil
}
