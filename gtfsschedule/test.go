package gtfsschedule

import (
	"fmt"
	"testing"
)

func TestParseTransfer(t *testing.T) {
	// agency.txt をテスト用のデータファイルとして使用する
	filePath := "transfers.txt"
	agencies, err := ParseTransfer(filePath)
	if err != nil {
		t.Fatalf("failed to parse CSV: %v", err)
	}

	fmt.Println(agencies)

	// 結果の検証
	if len(agencies) == 0 {
		t.Errorf("expected agencies to be non-empty, but got empty")
	}
}
