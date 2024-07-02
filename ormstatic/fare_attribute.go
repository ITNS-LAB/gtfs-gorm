package ormstatic

import "database/sql"

type FareAttribute struct {
	FareId           string        `gorm:"primaryKey;index;not null"`
	Price            float64       `gorm:"index;not null"`
	CurrencyType     string        `gorm:"primaryKey;not null"`
	PaymentMethod    int16         `gorm:"not null"`
	Transfers        sql.NullInt16 //必須の項目だが、空の場合があるため、nullを許容
	AgencyId         sql.NullString
	TransferDuration sql.NullInt32
	//FareRule         FareRule `gorm:"foreignKey:FareId"`
}

func (FareAttribute) TableName() string {
	return "fare_attributes"
}

func (f FareAttribute) GetFareId() any {
	return f.FareId
}

func (f FareAttribute) GetPrice() any {
	return f.Price
}

func (f FareAttribute) GetCurrencyType() any {
	return f.CurrencyType
}

func (f FareAttribute) GetPaymentMethod() any {
	return f.PaymentMethod
}

func (f FareAttribute) GetTransfers() any {
	if f.Transfers.Valid {
		return f.Transfers.Int16
	}
	return nil
}

func (f FareAttribute) GetAgencyId() any {
	if f.AgencyId.Valid {
		return f.AgencyId.String
	}
	return nil
}

func (f FareAttribute) GetTransferDuration() any {
	if f.TransferDuration.Valid {
		return f.TransferDuration.Int32
	}
	return nil
}
