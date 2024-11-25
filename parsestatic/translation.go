package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/gtfsjp"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
)

func ParseTranslations(path string) ([]gtfsjp.Translation, error) {
	var translations []gtfsjp.Translation
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

		translations = append(translations, gtfsjp.Translation{
			Tablename:   util.IsBlank(df.GetElement("table_name")),
			FieldName:   util.IsBlank(df.GetElement("field_name")),
			Language:    util.IsBlank(df.GetElement("language")),
			Translation: util.IsBlank(df.GetElement("translation")),
			RecordId:    util.IsBlank(df.GetElement("record_id")),
			RecordSubId: util.IsBlank(df.GetElement("record_sub_id")),
			FieldValue:  util.IsBlank(df.GetElement("field_value")),
		})
	}
	return translations, nil
}
