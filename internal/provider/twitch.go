package provider

import (
	"cli-stream-chat/internal"
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

func (t *Twitch) Listen(ctx context.Context, out chan internal.Message) error {
	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		userId, _ := strconv.Atoi(message.User.ID)
		out <- internal.Message{
			UserId:   userId,
			Nickname: message.User.DisplayName,
			Text:     message.Message,
			Platform: internal.TwitchPlatform,
		}
	})

	t.client.Join(t.channel)

	if err := t.client.Connect(); err != nil {
		return fmt.Errorf("twitch error: %w", err)
	}
	return nil
}
