package pipe

import (
	"cli-stream-chat/internal"
	"fmt"
	"os"
	"time"
)

type Log struct {
	Path string
}

func (s *Log) filename() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%d-%d-%d.log", year, month, day)
}

func (s *Log) fullpath() string {
	return s.Path + "/" + s.filename()
}

func (s *Log) Write(m internal.Message) error {
	// TODO: init log with file once, so we don't check file on every write
	file, err := getOrCreateFile(s.fullpath())
	defer file.Close()
	if err != nil {
		return fmt.Errorf("problem when open file: %w", err)
	}
	_, err = file.WriteString(m.FullText() + "\n")
	if err != nil {
		return fmt.Errorf("problem with write to file: %w", err)
	}
	return nil
}

func getOrCreateFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return file, err
	}
	return file, nil
}
