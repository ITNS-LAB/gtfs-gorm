package ormstatic

import "database/sql"

type Pathway struct {
	PathwayId            string `gorm:"primaryKey;index;not null"`
	FromStopId           string `gorm:"index;not null"`
	ToStopId             string `gorm:"index;not null"`
	PathwayMode          int16  `gorm:"index;not null"`
	IsBidirectional      int16  `gorm:"index;not null"`
	Length               sql.NullFloat64
	TraversalTime        sql.NullInt32
	StairCount           sql.NullInt32   //non null integer だが、データがないことを示すnullが必要なためNullInt型に指定
	MaxSlope             sql.NullFloat64 `gorm:"default:0"`
	MinWidth             sql.NullFloat64
	SignpostedAs         sql.NullString
	ReversedSignpostedAs sql.NullString
}

func (Pathway) TableName() string {
	return "pathways"
}

func (p Pathway) GetPathwayId() any {
	return p.PathwayId
}

func (p Pathway) GetFromStopId() any {
	return p.FromStopId
}

func (p Pathway) GetToStopId() any {
	return p.ToStopId
}

func (p Pathway) GetPathwayMode() any {
	return p.PathwayMode
}

func (p Pathway) GetIsBidirectional() any {
	return p.IsBidirectional
}

func (p Pathway) GetLength() any {
	if p.Length.Valid {
		return p.Length.Float64
	}
	return nil
}

func (p Pathway) GetTraversalTime() any {
	if p.TraversalTime.Valid {
		return p.TraversalTime.Int32
	}
	return nil
}

func (p Pathway) GetStairCount() any {
	if p.StairCount.Valid {
		return p.StairCount.Int32
	}
	return nil
}

func (p Pathway) GetMaxSlope() any {
	if p.MaxSlope.Valid {
		return p.MaxSlope.Float64
	}
	return nil
}

func (p Pathway) GetMinWidth() any {
	if p.MinWidth.Valid {
		return p.MinWidth.Float64
	}
	return nil
}

func (p Pathway) GetSignpostedAs() any {
	if p.SignpostedAs.Valid {
		return p.SignpostedAs.String
	}
	return nil
}

func (p Pathway) GetReversedSignpostedAs() any {
	if p.ReversedSignpostedAs.Valid {
		return p.ReversedSignpostedAs.String
	}
	return nil
}
