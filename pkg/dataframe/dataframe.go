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

func NewDataFrame() *DataFrame {
	return &DataFrame{Headers: make(map[string]int), Position: 0}
}

func (df *DataFrame) HasNext() bool {
	return df.Position < len(df.Rows)
}

func (df *DataFrame) Next() ([]string, error) {
	if df.HasNext() {
		row := df.Rows[df.Position]
		df.Position++
		return row, nil
	}
	return nil, errors.New("no more elements")
}

func (df *DataFrame) setHeader(header []string) {
	for i, v := range header {
		df.Headers[strings.TrimSpace(v)] = i
	}
}

func (df *DataFrame) AppendRow(row []string) {
	df.Rows = append(df.Rows, row)
}

func (df *DataFrame) HasHeader(s string) (bool, int) {
	i, exists := df.Headers[s]
	if !exists {
		return false, -1
	}
	return true, i
}

func (df *DataFrame) GetElement(s string) sql.NullString {
	exists, i := df.HasHeader(s)
	if exists {
		return sql.NullString{String: df.Rows[df.Position-1][i], Valid: true}
	}
	return sql.NullString{String: "", Valid: false}
}
