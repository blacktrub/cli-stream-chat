package internal

import (
	"cli-stream-chat/internal/sticker"
	"fmt"
)

type Platform string

const (
	TwitchPlatform  Platform = "TW"
	YoutubePlatform Platform = "YT"
)

type Message struct {
	UserId   int
	Nickname string
	Text     string
	Platform Platform
}

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s", m.Nickname, m.Text))
}

func (m *Message) ColorizedText() string {
	text := sticker.FindAndReplace(m.Text)
	return fmt.Sprintf("%s: %s", colorizer(m.Platform)(m.UserId, m.Nickname), text)
}

func colorizer(p Platform) func(int, string) string {
	crl := Colorizer{make(map[int]int)}
	switch p {
	case TwitchPlatform:
		return crl.Do
	case YoutubePlatform:
		return makeRed
	default:
		return func(i int, m string) string { return m }
	}
}

func makeRed(id int, m string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}
