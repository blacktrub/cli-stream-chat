package pipe

import (
	"cli-stream-chat/internal"
	"fmt"
	"os"
	"time"
)

type Log struct {
	f *os.File
}

func NewLog(f *os.File) *Log {
	return &Log{f: f}
}

func (s *Log) Write(m internal.Message) error {
	_, err := s.f.WriteString(m.FullText() + "\n")
	if err != nil {
		return fmt.Errorf("problem with write to file: %w", err)
	}
	return nil
}

// TODO: move it somewhere?
func GetFile(path string) (*os.File, error) {
	fullPath := getFileName(path)
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return file, err
	}
	return file, nil
}

func getFileName(path string) string {
	year, month, day := time.Now().Date()
	name := fmt.Sprintf("%d-%d-%d.log", year, month, day)
	return path + "/" + name
}
