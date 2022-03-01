package msg

import (
	"cli-stream-chat/pkg/platform"
	"fmt"
)

type Message struct {
	Nickname string
	Text     string
	platform.Platform
}

type MsgStream chan Message

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s", m.Nickname, m.Text))
}

func (m *Message) ColorizedText() string {
	colorize := getColorize(m.Platform)
	return fmt.Sprintf("%s: %s", colorize(m.Nickname), m.Text)
}

func getColorize(p platform.Platform) func(string) string {
	switch p.Name {
	case platform.Twitch:
		return makeBlue
	case platform.Youtube:
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
