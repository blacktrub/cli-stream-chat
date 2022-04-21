package internal

import (
	"fmt"
)

type Platform string

const (
	TwitchPlatform  Platform = "TW"
	YoutubePlatform Platform = "YT"
)

type Message struct {
	Nickname string
	Text     string
	Platform Platform
}

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s", m.Nickname, m.Text))
}

func (m *Message) ColorizedText() string {
	return fmt.Sprintf("%s: %s", colorizer(m.Platform)(m.Nickname), m.Text)
}

func colorizer(p Platform) func(string) string {
	switch p {
	case TwitchPlatform:
		return makeBlue
	case YoutubePlatform:
		return makeRed
	default:
		return func(m string) string { return m }
	}
}

func makeBlue(m string) string {
	return fmt.Sprintf("\033[1;34m%s\033[0m", m)
}

func makeRed(m string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}
