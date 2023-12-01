package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseFareAttributes(path string) []orm.FareAttribute {
	var fareAttributes []orm.FareAttribute
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fareAttributes = append(fareAttributes, orm.FareAttribute{
			FareId:           dataframe.IsBlank(df.GetElement("fare_id")),
			Price:            dataframe.ParseFloat64(df.GetElement("price")),
			CurrencyType:     dataframe.IsBlank(df.GetElement("currency_type")),
			PaymentMethod:    dataframe.ParseEnum(df.GetElement("payment_method")),
			Transfers:        dataframe.ParseEnum(df.GetElement("transfers")),
			AgencyId:         dataframe.IsBlank(df.GetElement("agency_id")),
			TransferDuration: dataframe.ParseInt(df.GetElement("transfer_duration")),
		})
	}
	return fareAttributes
}
