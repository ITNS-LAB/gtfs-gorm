package static

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/orm/static"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
)

func ParsePathways(path string) []static.Pathway {
	var pathways []static.Pathway
	df := dataframe.OpenCsv(path)
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		pathways = append(pathways, static.Pathway{
			PathwayId:            dataframe.IsBlank(df.GetElement("pathway_id")),
			FromStopId:           dataframe.IsBlank(df.GetElement("from_stop_id")),
			ToStopId:             dataframe.IsBlank(df.GetElement("to_stop_id")),
			PathwayMode:          dataframe.ParseEnum(df.GetElement("pathway_mode")),
			IsBidirectional:      dataframe.ParseEnum(df.GetElement("is_bidirectional")),
			Length:               dataframe.ParseFloat64(df.GetElement("length")),
			TraversalTime:        dataframe.ParseInt(df.GetElement("traversal_time")),
			StairCount:           dataframe.ParseInt(df.GetElement("stair_count")),
			MaxSlope:             dataframe.ParseFloat64(df.GetElement("max_slope")),
			MinWidth:             dataframe.ParseFloat64(df.GetElement("win_width")),
			SignpostedAs:         dataframe.IsBlank(df.GetElement("signposted_as")),
			ReversedSignpostedAs: dataframe.IsBlank(df.GetElement("reversed_signposted_as")),
		})
	}
	return pathways
}
