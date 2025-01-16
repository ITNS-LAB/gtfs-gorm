package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Translation struct {
	TableName   string  `gorm:"not null"`
	FieldName   string  `gorm:"not null"`
	Language    string  `gorm:"not null"`
	Translation string  `gorm:"not null"`
	RecordID    *string `json:"record_id,omitempty"`     // 翻訳するフィールドに対応するレコード ID
	RecordSubID *string `json:"record_sub_id,omitempty"` // サブ ID（必要な場合）
	FieldValue  *string `json:"field_value,omitempty"`   // 翻訳を適用する特定の値
}

func ParseTranslation(path string) ([]Translation, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open translation CSV: %w", err)
	}

	// データを解析してTranslation構造体のスライスを作成
	var translations []Translation
	for i := 0; i < len(df.Records); i++ {
		tableName, err := df.GetString(i, "table_name")
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
			// optional field, so it's okay if this is nil
		}

		recordSubId, err := df.GetStringPtr(i, "record_sub_id")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		fieldValue, err := df.GetStringPtr(i, "field_value")
		if err != nil {
			// optional field, so it's okay if this is nil
		}

		// Translation 構造体を作成しリストに追加
		translations = append(translations, Translation{
			TableName:   tableName,
			FieldName:   fieldName,
			Language:    language,
			Translation: translation,
			RecordID:    recordId,
			RecordSubID: recordSubId,
			FieldValue:  fieldValue,
		})
	}

	return translations, nil
}
