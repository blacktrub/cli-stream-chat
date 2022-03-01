package provider

import (
	"fmt"
	"sync"

	"cli-stream-chat/pkg/msg"
	"cli-stream-chat/pkg/pipe"
	"cli-stream-chat/pkg/platform"

	"github.com/abhinavxd/youtube-live-chat-downloader/v2"
)

type Youtube struct{}

func (y *Youtube) Listen(wg *sync.WaitGroup, streamLink string, pipes pipe.Pipes, ch msg.MsgStream) {
	continuation, cfg, error := YtChat.ParseInitialData(streamLink)
	if error != nil {
		fmt.Println("error youtube", error)
	}
	for {
		chat, newContinuation, error := YtChat.FetchContinuationChat(continuation, cfg)
		if error != nil {
			fmt.Println("error youtube", error)
			continue
		}
		continuation = newContinuation
		for _, m := range chat {
			ch <- msg.Message{
				Nickname: m.AuthorName,
				Text:     m.Message,
				Platform: platform.YoutubePlatform,
			}
		}
	}
}
