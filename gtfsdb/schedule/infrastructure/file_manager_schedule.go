package infrastructure

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsdb/schedule/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"os"
	"path/filepath"
)

type FileManagerSchedule struct{}

func (f FileManagerSchedule) Download(url, path string) error {
	if err := util.FetchFile(url, path); err != nil {
		return err
	}
	return nil
}

func (f FileManagerSchedule) Remove(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}
func (f FileManagerSchedule) UnZip(file, path string) (string, error) {
	dest, err := util.UnZip(file, path)
	if err != nil {
		return "", err
	}

	// サブディレクトリを探して、それを新しいルートにする
	entries, err := os.ReadDir(dest)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			return filepath.Join(dest, entry.Name()), nil
		}
	}

	// フォールバック：サブディレクトリがなければ元のdestを返す
	return dest, nil
}

func NewFileManagerRepository() repository.FileManagerRepository {
	return &FileManagerSchedule{}
}
