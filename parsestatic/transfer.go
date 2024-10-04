package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
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
			fmt.Println("Error:", err)
			break
		}

		transfers = append(transfers, ormstatic.Transfer{
			FromStopId:      util.IsBlank(df.GetElement("from_stop_id")),
			ToStopId:        util.IsBlank(df.GetElement("to_stop_id")),
			TransferType:    util.ParseEnum(df.GetElement("transfer_type")),
			MinTransferTime: util.ParseInt(df.GetElement("min_transfer_time")),
		})
	}
	return transfers, nil
}
