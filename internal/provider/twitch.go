package provider

import (
	"cli-stream-chat/internal/color"
	"cli-stream-chat/internal/msg"
	"cli-stream-chat/internal/sticker"
	"context"
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v3"
)

type Twitch struct {
	client    *twitch.Client
	channel   string // twitch channel name
	colorizer msg.Colorizer
}

func NewTwitchProvider(channel string) *Twitch {
	return &Twitch{
		channel:   channel,
		client:    twitch.NewAnonymousClient(),
		colorizer: color.NewPretty(),
	}
}

func (t *Twitch) Listen(ctx context.Context, out chan msg.Message) error {
	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// TODO: move it somewhere
		emotes := sticker.TwitchEmotes{}
		for i := 0; i < len(message.Emotes); i++ {
			e := message.Emotes[i]
			emotes = append(emotes, sticker.TwitchEmote{ID: e.ID, Name: e.Name})
		}

		userId, _ := strconv.Atoi(message.User.ID)
		out <- *msg.NewTwitch(
			userId,
			message.User.DisplayName,
			message.Message,
			message.User.Badges,
			message.RoomID,
			emotes,
			t.colorizer,
		)
	})

	t.client.Join(t.channel)

	if err := t.client.Connect(); err != nil {
		return fmt.Errorf("twitch error: %w", err)
	}
	return nil
}
