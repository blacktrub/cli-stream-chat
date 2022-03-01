package provider

import (
	"sync"

	"cli-stream-chat/pkg/msg"
	"cli-stream-chat/pkg/pipe"
	"cli-stream-chat/pkg/platform"

	"github.com/gempir/go-twitch-irc/v3"
)

type Twitch struct{}

func (t *Twitch) Listen(wg *sync.WaitGroup, channelName string, pipes pipe.Pipes, ch msg.MsgStream) {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		ch <- msg.Message{
			Nickname: message.User.DisplayName,
			Text:     message.Message,
			Platform: platform.TwitchPlatform,
		}
	})

	client.Join(channelName)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
