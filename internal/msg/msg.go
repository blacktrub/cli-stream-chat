package msg

import (
	"cli-stream-chat/internal/badge"
	"cli-stream-chat/internal/color"
	"cli-stream-chat/internal/sticker"
	"fmt"
)

type Platform string

type Colorizer interface {
	Colorize(int, string) string
}

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
	Emotes    sticker.TwitchEmotes
	Colorizer Colorizer
}

func NewTwitch(
	userId int,
	nickname string,
	text string,
	badges map[string]int,
	boId string,
	emotes sticker.TwitchEmotes,
	colorizer Colorizer,
) *Message {
	return &Message{
		UserId:        userId,
		Nickname:      nickname,
		Text:          text,
		Platform:      TwitchPlatform,
		Badges:        badges,
		BroadcasterId: boId,
		Emotes:        emotes,
		Colorizer:     colorizer,
	}
}

func NewYoutube(
	nickname string,
	text string,
) *Message {
	return &Message{
		Nickname:  nickname,
		Text:      text,
		Platform:  YoutubePlatform,
		Colorizer: color.MakeRed{},
	}
}

func (m *Message) FullText() string {
	return fmt.Sprintf(fmt.Sprintf("%s: %s", m.Nickname, m.Text))
}

func (m *Message) PrettyText() string {
	text := sticker.FindAndReplace(m.Text, m.Emotes, m.BroadcasterId)
	badges := badge.Show(m.Badges)
	nickname := m.Colorizer.Colorize(m.UserId, m.Nickname)
	// TODO: maybe we need some space between badges and a nickname
	return fmt.Sprintf("%s%s: %s", badges, nickname, text)
}
