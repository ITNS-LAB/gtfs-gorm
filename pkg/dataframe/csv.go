package dataframe

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"os"
)

func OpenCsv(f string) DataFrame {
	// ファイルをオープン
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// CSVリーダーを作成 BOMの対策もする
	fallback := unicode.UTF8.NewDecoder()
	reader := csv.NewReader(transform.NewReader(file, unicode.BOMOverride(fallback)))

	// ヘッダーを読み込む
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading Headers:", err)
	}

	// データフレームの作成とヘッダーの挿入
	df := NewDataFrame()
	df = df.setHeader(headers)

	for {
		record, err := reader.Read()
		if err != nil {
			break // ファイルの終わりに達したらループを終了
		}

		df = df.AppendRow(record)
	}
	return df
}
