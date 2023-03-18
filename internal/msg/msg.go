package msg

import (
	"cli-stream-chat/internal"
	"cli-stream-chat/internal/badge"
	"cli-stream-chat/internal/sticker"
	"fmt"
)

type Platform string

const (
	// TODO: why is it here?
	TwitchPlatform  Platform = "TW"
	YoutubePlatform Platform = "YT"
)

// TODO: create an interface and two implementations
// TODO: for twitch and youtube
type Message struct {
	UserId        int
	Nickname      string
	Text          string
	Platform      Platform
	Badges        map[string]int
	BroadcasterId string
	// TODO: do not use own type for a slice there
	Emotes sticker.TwitchEmotes
}

func NewTwitch(
	userId int,
	nickname string,
	text string,
	badges map[string]int,
	boId string,
	emotes sticker.TwitchEmotes,
) *Message {
	return &Message{
		UserId:        userId,
		Nickname:      nickname,
		Text:          text,
		Platform:      TwitchPlatform,
		Badges:        badges,
		BroadcasterId: boId,
		Emotes:        emotes,
	}
}

func NewYoutube(
	nickname string,
	text string,
) *Message {
	return &Message{
		Nickname: nickname,
		Text:     text,
		Platform: YoutubePlatform,
	}
}

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s", m.Nickname, m.Text))
}

func (m *Message) PrettyText() string {
	text := sticker.FindAndReplace(m.Text, m.Emotes, m.BroadcasterId)
	badges := badge.Show(m.Badges)
	// TODO: maybe we need some space between badges and a nickname
	return fmt.Sprintf("%s%s: %s", badges, colorizer(m.Platform)(m.UserId, m.Nickname), text)
}

func colorizer(p Platform) func(int, string) string {
	switch p {
	case TwitchPlatform:
		return internal.Crl.Do
	case YoutubePlatform:
		return makeRed
	default:
		return func(i int, m string) string { return m }
	}
}

func makeRed(id int, m string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}
