package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func FetchFile(url string, saveName string) error {
	// HTTPリクエストを作成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("リクエスト作成エラー: %w", err)
	}

	// HTTPクライアントを作成
	client := &http.Client{}

	// リクエストを送信し、レスポンスを取得
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("リクエスト送信エラー: %w", err)
	}
	defer resp.Body.Close()

	// レスポンスステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTPエラー: %d", resp.StatusCode)
	}

	// ディレクトリを作成
	dir := filepath.Dir(saveName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("ディレクトリ作成エラー: %w", err)
	}

	// ファイルを作成
	file, err := os.Create(saveName)
	if err != nil {
		return fmt.Errorf("ファイル作成エラー: %w", err)
	}
	defer file.Close()

	// レスポンスボディをファイルに保存
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("ファイル保存エラー: %w", err)
	}

	return nil
}
