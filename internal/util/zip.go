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
	if err := os.MkdirAll("temp", 0755); err != nil {
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
		if err := saveUnZipFiles("temp", f); err != nil {
			return "", err
		}
	}

	return "temp", nil
}

func saveUnZipFiles(destDir string, f *zip.File) error {
	// __MACOSX フォルダはスキップ
	if strings.Contains(f.Name, "__MACOSX") {
		return nil
	}

	relPath := filepath.Base(f.Name) // フラット展開
	destPath := filepath.Join(destDir, relPath)

	// ZIPスリップ攻撃対策
	if !strings.HasPrefix(destPath, filepath.Clean(destDir)+string(os.PathSeparator)) {
		return fmt.Errorf("不正なパス: %s", destPath)
	}

	// 親ディレクトリ作成
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, rc); err != nil {
		return err
	}

	return nil
}
