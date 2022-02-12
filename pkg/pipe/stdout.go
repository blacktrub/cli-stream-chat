package pipe

import "fmt"

type Stdout struct{}

func (s *Stdout) Write(msg Message) {
	fmt.Println(msg.ColorizedText())
}
