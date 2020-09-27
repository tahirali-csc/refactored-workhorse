package logstorage

import (
	"context"
	"os"
)

type FileLogStore struct {
	file *os.File
}

func NewFileLogStore(file *os.File) LogStore {
	return &FileLogStore{
		file: file,
	}
}

func (fl *FileLogStore) Write(ctx context.Context, stepId int64, line []byte) error {
	_, err := fl.file.Write(line)
	return err
}
