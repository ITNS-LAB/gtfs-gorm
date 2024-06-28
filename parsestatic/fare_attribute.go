package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseFareAttributes(path string) ([]ormstatic.FareAttribute, error) {
	var fareAttributes []ormstatic.FareAttribute
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return fareAttributes, err
	}

	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			return []ormstatic.FareAttribute{}, err
		}

		fareId, err := dataframe.ParseString(df.GetElement("fare_id"))
		if err != nil {
			return []ormstatic.FareAttribute{}, err
		}

		price, err := dataframe.ParseFloat64(df.GetElement("price"))
		if err != nil {
			return []ormstatic.FareAttribute{}, err
		}

		currencyType, err := dataframe.ParseString(df.GetElement("currency_type"))
		if err != nil {
			return []ormstatic.FareAttribute{}, err
		}

		paymentMethod, err := dataframe.ParseInt16(df.GetElement("payment_method"))
		if err != nil {
			return []ormstatic.FareAttribute{}, err
		}

		transfers, err := dataframe.ParseNullInt16(df.GetElement("transfers"))
		transferDuration, err := dataframe.ParseNullInt32(df.GetElement("transfer_duration"))

		fareAttributes = append(fareAttributes, ormstatic.FareAttribute{
			FareId:           fareId,
			Price:            price,
			CurrencyType:     currencyType,
			PaymentMethod:    paymentMethod,
			Transfers:        transfers,
			AgencyId:         df.GetElement("agency_id"),
			TransferDuration: transferDuration,
		})
	}
	return fareAttributes, nil
}
