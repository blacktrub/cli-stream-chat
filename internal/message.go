package internal

import (
	"cli-stream-chat/internal/badge"
	"cli-stream-chat/internal/sticker"
	"fmt"
)

type Platform string

const (
	TwitchPlatform  Platform = "TW"
	YoutubePlatform Platform = "YT"
)

type Message struct {
	UserId        int
	Nickname      string
	Text          string
	Platform      Platform
	Badges        map[string]int
	BroadcasterId string
}

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s", m.Nickname, m.Text))
}

func (m *Message) ColorizedText() string {
	text := sticker.FindAndReplace(m.Text, m.BroadcasterId)
	badges := badge.Show(m.Badges)
	// TODO: maybe we need some space between badges and a nickname
	return fmt.Sprintf("%s%s: %s", badges, colorizer(m.Platform)(m.UserId, m.Nickname), text)
}

func colorizer(p Platform) func(int, string) string {
	switch p {
	case TwitchPlatform:
		return Crl.Do
	case YoutubePlatform:
		return makeRed
	default:
		return func(i int, m string) string { return m }
	}
}

func makeRed(id int, m string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}
