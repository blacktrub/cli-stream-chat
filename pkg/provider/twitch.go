package provider

import (
	"sync"

	"cli-stream-chat/pkg/pipe"

	"github.com/gempir/go-twitch-irc/v3"
)

type Twitch struct{}

func (t *Twitch) Listen(wg *sync.WaitGroup, channelName string, pipes pipe.Pipes, ch pipe.MsgStream) {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		ch <- pipe.Message{Nickname: message.User.DisplayName, Text: message.Message, Platform: pipe.TwitchPlatform}
	})

	client.Join(channelName)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
