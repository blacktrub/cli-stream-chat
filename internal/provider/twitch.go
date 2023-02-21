package provider

import (
	int "cli-stream-chat/internal"
	"cli-stream-chat/internal/sticker"
	"context"
	"fmt"
	"strconv"

	"github.com/gempir/go-twitch-irc/v3"
)

type Twitch struct {
	client  *twitch.Client
	channel string // twitch channel name
}

func NewTwitchProvider(channel string) *Twitch {
	return &Twitch{
		channel: channel,
		client:  twitch.NewAnonymousClient(),
	}
}

func (t *Twitch) Listen(ctx context.Context, out chan int.Message) error {
	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// TODO: move it somewhere
		emotes := sticker.TwitchEmotes{}
		for i := 0; i < len(message.Emotes); i++ {
			e := message.Emotes[i]
			emotes = append(emotes, sticker.TwitchEmote{ID: e.ID, Name: e.Name})
		}

		userId, _ := strconv.Atoi(message.User.ID)
		out <- int.Message{
			UserId:        userId,
			Badges:        message.User.Badges,
			Nickname:      message.User.DisplayName,
			Text:          message.Message,
			Platform:      int.TwitchPlatform,
			BroadcasterId: message.RoomID,
			Emotes:        emotes,
		}
	})

	t.client.Join(t.channel)

	if err := t.client.Connect(); err != nil {
		return fmt.Errorf("twitch error: %w", err)
	}
	return nil
}
