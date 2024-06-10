package dataframe

import (
	"database/sql"
	"errors"
	"strings"
)

type DataFrame struct {
	Headers  map[string]int
	Rows     [][]string
	Position int
}

func NewDataFrame() DataFrame {
	df := DataFrame{Headers: make(map[string]int), Position: 0}
	return df
}

func (df *DataFrame) HasNext() bool {
	return df.Position < len(df.Rows)
}

func (df *DataFrame) Next() ([]string, error) {
	if df.HasNext() {
		rows := df.Rows[df.Position]
		df.Position++
		return rows, nil
	}
	return nil, errors.New("no more elements")
}

func (df *DataFrame) setHeader(header []string) DataFrame {
	tempDf := *df
	for i, v := range header {
		tempDf.Headers[strings.TrimSpace(v)] = i
	}
	return tempDf
}

func (df *DataFrame) AppendRow(row []string) DataFrame {
	tempDf := *df
	tempDf.Rows = append(tempDf.Rows, [][]string{row}...)
	return tempDf
}

func (df *DataFrame) HasHeader(s string) (bool, int) {
	i, isExists := df.Headers[s]
	if isExists == false {
		i = -1
	}
	return isExists, i
}

func (df *DataFrame) GetElement(s string) sql.NullString {
	exists, i := df.HasHeader(s)
	if exists {
		return sql.NullString{String: df.Rows[df.Position-1][i], Valid: true}
	} else {
		return sql.NullString{String: "", Valid: false}
	}
}
