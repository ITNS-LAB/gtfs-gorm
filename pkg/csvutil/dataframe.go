package csvutil

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"gorm.io/datatypes"
	"os"
	"strconv"
	"time"
)

const (
	dateFormat = "20060102" // 日付フォーマット
	timeFormat = "15:04:05" // 時刻フォーマット
)

// DataFrame はヘッダー情報と行データを保持する構造体です。
type DataFrame struct {
	Headers map[string]int
	Records [][]string
}

// OpenCSV はBOM付きのCSVファイルを開き、ヘッダーとデータを読み込みます。
func OpenCSV(filePath string) (*DataFrame, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// BOMを除去するリーダーを作成
	reader := csv.NewReader(transform.NewReader(file, unicode.BOMOverride(unicode.UTF8.NewDecoder())))

	// ヘッダーを取得
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read Headers: %w", err)
	}

	// データを読み込む
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	// ヘッダーをインデックスマップに変換
	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	return &DataFrame{
		Headers: headerMap,
		Records: records,
	}, nil
}

// getValue は指定されたカラムと行番号から値を取得します。
func (df *DataFrame) getValue(rowIndex int, columnName string) (string, error) {
	idx, exists := df.Headers[columnName]
	if !exists {
		return "", fmt.Errorf("column '%s' does not exist", columnName)
	}

	if rowIndex < 0 || rowIndex >= len(df.Records) {
		return "", fmt.Errorf("row index %d out of range", rowIndex)
	}

	return df.Records[rowIndex][idx], nil
}

// GetString は指定されたカラムの値を文字列として取得します。
func (df *DataFrame) GetString(rowIndex int, columnName string) (string, error) {
	return df.getValue(rowIndex, columnName)
}

// GetStringPtr は指定されたカラムの値をポインタ型の文字列として取得します。
func (df *DataFrame) GetStringPtr(rowIndex int, columnName string) (*string, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return nil, nil
	}
	if value == "" {
		return nil, nil
	}
	return &value, nil
}

// GetInt は指定されたカラムの値を整数として取得します。
func (df *DataFrame) GetInt(rowIndex int, columnName string) (int, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return 0, err
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	return intValue, nil
}

// GetIntPtr は指定されたカラムの値をポインタ型の整数として取得します。
func (df *DataFrame) GetIntPtr(rowIndex int, columnName string) (*int, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return nil, nil
	}
	if value == "" {
		return nil, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse int in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	return &intValue, nil
}

// GetFloat は指定されたカラムの値を浮動小数点数として取得します。
func (df *DataFrame) GetFloat(rowIndex int, columnName string) (float64, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return 0.0, err
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0, fmt.Errorf("failed to parse float in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	return floatValue, nil
}

// GetFloatPtr は指定されたカラムの値をポインタ型の浮動小数点数として取得します。
func (df *DataFrame) GetFloatPtr(rowIndex int, columnName string) (*float64, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return nil, nil
	}
	if value == "" {
		return nil, nil
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse float in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	return &floatValue, nil
}

// GetDate は指定されたカラムの値を日付型として取得します。
func (df *DataFrame) GetDate(rowIndex int, columnName string) (datatypes.Date, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return datatypes.Date{}, err
	}
	parsed, err := time.Parse(dateFormat, value)
	if err != nil {
		return datatypes.Date{}, fmt.Errorf("failed to parse date in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	return datatypes.Date(parsed), nil
}

// GetDatePtr は指定されたカラムの値をポインタ型の日付として取得します。
func (df *DataFrame) GetDatePtr(rowIndex int, columnName string) (*datatypes.Date, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return nil, nil
	}
	if value == "" {
		return nil, nil
	}
	parsed, err := time.Parse(dateFormat, value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	t := datatypes.Date(parsed)
	return &t, nil
}

// GetTime は指定されたカラムの値を時刻型として取得します。
func (df *DataFrame) GetTime(rowIndex int, columnName string) (datatypes.Time, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return datatypes.NewTime(0, 0, 0, 0), err
	}
	parsed, err := time.Parse(timeFormat, value)
	if err != nil {
		return datatypes.NewTime(0, 0, 0, 0), fmt.Errorf("failed to parse time in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	return datatypes.NewTime(parsed.Hour(), parsed.Minute(), parsed.Second(), 0), nil
}

// GetTimePtr は指定されたカラムの値をポインタ型の時刻として取得します。
func (df *DataFrame) GetTimePtr(rowIndex int, columnName string) (*datatypes.Time, error) {
	value, err := df.getValue(rowIndex, columnName)
	if err != nil {
		return nil, nil
	}
	if value == "" {
		return nil, nil
	}
	parsed, err := time.Parse(timeFormat, value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse time in column '%s' at row %d: %w", columnName, rowIndex, err)
	}
	t := datatypes.NewTime(parsed.Hour(), parsed.Minute(), parsed.Second(), 0)
	return &t, nil
}

//GetBoolPtrは指定されたカラムの値をポインタ型のプリミティブ型として取得します。
