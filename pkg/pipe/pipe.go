package pipe

import "fmt"

// TODO: maybe platform must be in separate package
type Platform struct {
	Name string
}

type Message struct {
	Nickname string
	Text     string
	Platform
}

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s\n", m.Nickname, m.Text))
}

type Pipe interface {
	Write(Message)
}

type Pipes []Pipe

const Twitch string = "TW"
const Youtube string = "YT"

func WriteAll(pipes Pipes, msg Message) {
	for i := 0; i < len(pipes); i++ {
		pipes[i].Write(msg)
	}
}
