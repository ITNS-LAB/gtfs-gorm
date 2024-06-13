package parsestatic

import (
	"github.com/ITNS-LAB/gtfs-gorm/ormstatic"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/dataframe"
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
			return []ormstatic.Pathway{}, err
		}

		pathwayId, err := dataframe.ParseString(df.GetElement("pathway_id"))
		if err != nil {
			return []ormstatic.Pathway{}, err
		}

		fromStopId, err := dataframe.ParseString(df.GetElement("from_stop_id"))
		if err != nil {
			return []ormstatic.Pathway{}, err
		}

		toStopId, err := dataframe.ParseString(df.GetElement("to_stop_id"))
		if err != nil {
			return []ormstatic.Pathway{}, err
		}

		pathwayMode, err := dataframe.ParseInt16(df.GetElement("pathway_mode"))
		if err != nil {
			return []ormstatic.Pathway{}, err
		}

		isBidirectional, err := dataframe.ParseInt16(df.GetElement("is_bidirectional"))
		if err != nil {
			return []ormstatic.Pathway{}, err
		}

		length, err := dataframe.ParseNullFloat64(df.GetElement("length"))
		traversalTime, err := dataframe.ParseNullInt32(df.GetElement("traversal_time"))
		stairCount, err := dataframe.ParseNullInt32(df.GetElement("stair_count"))
		maxSlope, err := dataframe.ParseNullFloat64(df.GetElement("max_slope"))
		minWidth, err := dataframe.ParseNullFloat64(df.GetElement("win_width"))

		pathways = append(pathways, ormstatic.Pathway{
			PathwayId:            pathwayId,
			FromStopId:           fromStopId,
			ToStopId:             toStopId,
			PathwayMode:          pathwayMode,
			IsBidirectional:      isBidirectional,
			Length:               length,
			TraversalTime:        traversalTime,
			StairCount:           stairCount,
			MaxSlope:             maxSlope,
			MinWidth:             minWidth,
			SignpostedAs:         df.GetElement("signposted_as"),
			ReversedSignpostedAs: df.GetElement("reversed_signposted_as"),
		})
	}
	return pathways, nil
}
