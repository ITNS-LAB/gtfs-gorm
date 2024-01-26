package parsestatic

import (
	"fmt"
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
			fmt.Println("Error:", err)
			break
		}

		translations = append(translations, ormstatic.Translation{
			Tablename:   dataframe.IsBlank(df.GetElement("table_name")),
			FieldName:   dataframe.IsBlank(df.GetElement("field_name")),
			Language:    dataframe.IsBlank(df.GetElement("language")),
			Translation: dataframe.IsBlank(df.GetElement("translation")),
			RecordId:    dataframe.IsBlank(df.GetElement("record_id")),
			RecordSubId: dataframe.IsBlank(df.GetElement("record_sub_id")),
			FieldValue:  dataframe.IsBlank(df.GetElement("field_value")),
		})
	}
	return translations, nil
}
