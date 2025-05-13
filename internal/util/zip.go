package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnZip(src, dest string) (string, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return "", err
	}
	defer r.Close()

	// 展開先ディレクトリを作成
	if err := os.MkdirAll(dest, 0755); err != nil {
		return "", err
	}

	// 全ファイル展開
	for _, f := range r.File {
		// __MACOSX フォルダ内のファイルをスキップ
		if strings.Contains(f.Name, "__MACOSX") {
			continue
		}
		if f.Mode().IsDir() {
			continue
		}
		if err := saveUnZipFiles(dest, f); err != nil {
			return "", err
		}
	}

	return dest, nil
}

func saveUnZipFiles(destDir string, f *zip.File) error {
	relPath := f.Name
	parts := strings.SplitN(relPath, string(os.PathSeparator), 2)
	if len(parts) == 2 && !strings.Contains(parts[0], "__MACOSX") {
		relPath = parts[1]
	}

	// 展開先のパスを設定する
	destPath := filepath.Join(destDir, relPath)

	// ZIPスリップ攻撃対策: 展開先のパスが期待されるディレクトリ内にあることを確認
	if !strings.HasPrefix(destPath, filepath.Clean(destDir)+string(os.PathSeparator)) {
		return fmt.Errorf("不正なパス: %s", destPath)
	}

	// 子孫ディレクトリがあれば作成する
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	// Zipファイルを開く
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// 展開先ファイルを作成する
	destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 展開先ファイルに書き込む
	if _, err := io.Copy(destFile, rc); err != nil {
		return err
	}

	return nil
}
