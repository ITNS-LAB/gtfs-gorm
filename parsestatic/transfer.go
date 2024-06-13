package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParseTransfers(path string) ([]ormstatic.Transfer, error) {
	var transfers []ormstatic.Transfer
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return transfers, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			return []ormstatic.Transfer{}, err
		}

		fromStopId, err := dataframe.ParseString(df.GetElement("from_stop_id"))
		if err != nil {
			return []ormstatic.Transfer{}, err
		}

		toStopId, err := dataframe.ParseString(df.GetElement("to_stop_id"))
		if err != nil {
			return []ormstatic.Transfer{}, err
		}

		transferType, err := dataframe.ParseInt16(df.GetElement("transfer_type"))
		minTransferTime, err := dataframe.ParseNullInt32(df.GetElement("min_transfer_time"))

		transfers = append(transfers, ormstatic.Transfer{
			FromStopId:      fromStopId,
			ToStopId:        toStopId,
			FromRouteId:     df.GetElement("from_route_id"),
			ToRouteId:       df.GetElement("to_route_id"),
			FromTripId:      df.GetElement("from_trip_id"),
			ToTripId:        df.GetElement("to_trip_id"),
			TransferType:    transferType,
			MinTransferTime: minTransferTime,
		})
	}
	return transfers, nil
}
