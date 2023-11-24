package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseTranslations(path string) []static.Translation {
	var translations []static.Translation
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		translations = append(translations, static.Translation{
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
