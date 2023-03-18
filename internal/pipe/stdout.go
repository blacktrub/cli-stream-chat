package pipe

import (
	"cli-stream-chat/internal/msg"
	"fmt"
)

type Stdout struct{}

func (s *Stdout) Write(m msg.Message) error {
	fmt.Println(m.PrettyText())
	return nil
}
