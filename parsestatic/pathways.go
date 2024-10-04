package parsestatic

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/internal/dataframe"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
)

func ParsePathways(path string) ([]ormstatic.Pathway, error) {
	var pathways []ormstatic.Pathway
	df, err := dataframe.OpenCsv(path)
	if err != nil {
		return pathways, err
	}
	for df.HasNext() {
		_, err := df.Next()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		pathways = append(pathways, ormstatic.Pathway{
			PathwayId:            util.IsBlank(df.GetElement("pathway_id")),
			FromStopId:           util.IsBlank(df.GetElement("from_stop_id")),
			ToStopId:             util.IsBlank(df.GetElement("to_stop_id")),
			PathwayMode:          util.ParseEnum(df.GetElement("pathway_mode")),
			IsBidirectional:      util.ParseEnum(df.GetElement("is_bidirectional")),
			Length:               util.ParseFloat64(df.GetElement("length")),
			TraversalTime:        util.ParseInt(df.GetElement("traversal_time")),
			StairCount:           util.ParseInt(df.GetElement("stair_count")),
			MaxSlope:             util.ParseFloat64(df.GetElement("max_slope")),
			MinWidth:             util.ParseFloat64(df.GetElement("win_width")),
			SignpostedAs:         util.IsBlank(df.GetElement("signposted_as")),
			ReversedSignpostedAs: util.IsBlank(df.GetElement("reversed_signposted_as")),
		})
	}
	return pathways, nil
}
