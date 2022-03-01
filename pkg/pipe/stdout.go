package pipe

import (
	"cli-stream-chat/pkg/msg"
	"fmt"
)

type Stdout struct{}

func (s *Stdout) Write(msg msg.Message) {
	fmt.Println(msg.ColorizedText())
}
