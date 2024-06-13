package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseTranslations(path string) ([]ormstatic.Translation, error) {
	var translations []ormstatic.Translation
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return translations, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			return []ormstatic.Translation{}, err
		}

		tableName, err := dataframe.ParseString(df.GetElement("table_name"))
		if err != nil {
			return []ormstatic.Translation{}, err
		}

		fieldName, err := dataframe.ParseString(df.GetElement("field_name"))
		if err != nil {
			return []ormstatic.Translation{}, err
		}

		language, err := dataframe.ParseString(df.GetElement("language"))
		if err != nil {
			return []ormstatic.Translation{}, err
		}

		translation, err := dataframe.ParseString(df.GetElement("translation"))
		if err != nil {
			return []ormstatic.Translation{}, err
		}

		translations = append(translations, ormstatic.Translation{
			Tablename:   tableName,
			FieldName:   fieldName,
			Language:    language,
			Translation: translation,
			RecordId:    df.GetElement("record_id"),
			RecordSubId: df.GetElement("record_sub_id"),
			FieldValue:  df.GetElement("field_value"),
		})
	}
	return translations, nil
}
