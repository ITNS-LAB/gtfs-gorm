package parser

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/static/orm"
)

func ParseTransfers(path string) []orm.Transfer {
	var transfers []orm.Transfer
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		transfers = append(transfers, orm.Transfer{
			FromStopId:      dataframe.IsBlank(df.GetElement("from_stop_id")),
			ToStopId:        dataframe.IsBlank(df.GetElement("to_stop_id")),
			TransferType:    dataframe.ParseEnum(df.GetElement("transfer_type")),
			MinTransferTime: dataframe.ParseInt(df.GetElement("min_transfer_time")),
		})
	}
	return transfers
}
