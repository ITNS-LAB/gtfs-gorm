package repository

type FileManagerRepository interface {
	Download(url, path string) error
	Remove(path string) error
	UnZip(file, path string) (string, error)
}
