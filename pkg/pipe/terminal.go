package pipe

import (
	"cli-stream-chat/pkg/msg"
	"fmt"
	"os"
)

type Device struct {
	Path string
}

func (s *Device) Write(msg msg.Message) {
	device, err := os.OpenFile(s.Path, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("problem when open device", err)
		return
	}
	defer device.Close()
	_, err = device.WriteString(msg.PrettyText() + "\n")
	if err != nil {
		fmt.Println("problem with write to device", err)
		return
	}
}
