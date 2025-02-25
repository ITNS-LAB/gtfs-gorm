package infrastructure

import (
	"github.com/ITNS-LAB/gtfs-gorm/gtfsscheduledb/domain/repository"
	"github.com/ITNS-LAB/gtfs-gorm/internal/util"
	"os"
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
	return dest, nil
}

func NewFileManagerRepository() repository.FileManagerRepository {
	return &FileManagerSchedule{}
}
