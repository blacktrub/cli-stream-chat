package pipe

import (
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

func (s *Log) Write(msg Message) {
	file, err := getOrCreateFile(s.fullpath())
	defer file.Close()
	if err != nil {
		fmt.Println("problem when open file", err)
		return
	}
	_, err = file.WriteString(msg.FullText() + "\n")
	if err != nil {
		fmt.Println("problem with write to file", err)
		return
	}
}

func getOrCreateFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return file, err
	}
	return file, nil
}
