package pipe

import (
	"cli-stream-chat/internal"
	"fmt"
)

type Stdout struct{}

func (s *Stdout) Write(m internal.Message) error {
	fmt.Println(m.ColorizedText())
	return nil
}
