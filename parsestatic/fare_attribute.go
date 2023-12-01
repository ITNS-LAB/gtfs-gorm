package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseFareAttributes(path string) []ormstatic.FareAttribute {
	var fareAttributes []ormstatic.FareAttribute
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fareAttributes = append(fareAttributes, ormstatic.FareAttribute{
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
