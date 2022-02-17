package pipe

import "fmt"

// TODO: maybe platform must be in separate package
type Platform struct {
	Name string
}

// TODO: maybe message must be in separate package
type Message struct {
	Nickname string
	Text     string
	Platform
}

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s", m.Nickname, m.Text))
}

func (m *Message) ColorizedText() string {
	colorize := getColorize(m.Platform)
	return fmt.Sprintf("%s: %s", colorize(m.Nickname), m.Text)
}

func getColorize(p Platform) func(string) string {
	switch p.Name {
	case Twitch:
		return makeBlue
	case Youtube:
		return makeRed
	default:
		return withoutColor
	}
}

func makeBlue(m string) string {
	return fmt.Sprintf("\033[1;34m%s\033[0m", m)
}

func makeRed(m string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}

func withoutColor(m string) string {
	return m
}

type Pipe interface {
	Write(Message)
}

type Pipes []Pipe

const Twitch string = "TW"
const Youtube string = "YT"

// TODO: I think I can use channel to do this work
func WriteAll(pipes Pipes, msg Message) {
	for i := 0; i < len(pipes); i++ {
		pipes[i].Write(msg)
	}
}
