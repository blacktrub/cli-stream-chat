package pipe

import (
	"cli-stream-chat/internal/msg"
	"fmt"
	"os"
)

type Device struct {
	Path string
}

func (s *Device) Write(m msg.Message) error {
	device, err := os.OpenFile(s.Path, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("problem when open device: %w", err)
	}
	defer device.Close()
	_, err = device.WriteString(m.PrettyText() + "\n")
	if err != nil {
		return fmt.Errorf("problem with write to device: %w", err)
	}
	return nil
}
