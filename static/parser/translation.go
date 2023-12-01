package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseTranslations(path string) []orm.Translation {
	var translations []orm.Translation
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		translations = append(translations, orm.Translation{
			Tablename:   dataframe.IsBlank(df.GetElement("table_name")),
			FieldName:   dataframe.IsBlank(df.GetElement("field_name")),
			Language:    dataframe.IsBlank(df.GetElement("language")),
			Translation: dataframe.IsBlank(df.GetElement("translation")),
			RecordId:    dataframe.IsBlank(df.GetElement("record_id")),
			RecordSubId: dataframe.IsBlank(df.GetElement("record_sub_id")),
			FieldValue:  dataframe.IsBlank(df.GetElement("field_value")),
		})
	}
	return translations
}
