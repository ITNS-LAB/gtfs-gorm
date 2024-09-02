package util

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func Sha256(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	// ハッシュのバイト配列を16進数の文字列に変換
	hashStr := hex.EncodeToString(hash[:])
	return hashStr, nil
}
