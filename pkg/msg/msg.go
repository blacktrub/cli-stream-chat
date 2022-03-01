package msg

import (
	"cli-stream-chat/pkg/platform"
	"cli-stream-chat/pkg/sticker"
	"fmt"
	"strings"
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

func (m *Message) PrettyText() string {
	// TODO: use only for Kitty terminal
	text := findAndReplaceStikers(m.Text)
	colorize := getColorize(m.Platform)
	return fmt.Sprintf("%s: %s", colorize(m.Nickname), text)
}

func findAndReplaceStikers(txt string) string {
	stickers := sticker.GetSupportedNames()
	for _, name := range stickers {
		if !strings.Contains(txt, name) {
			continue
		}
		buildedSticker := sticker.BuildKittyStiker(name)
		txt = strings.ReplaceAll(txt, name, buildedSticker)
	}
	return txt
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
