package provider

import (
	"fmt"
	"sync"

	"cli-stream-chat/pkg/pipe"

	"github.com/abhinavxd/youtube-live-chat-downloader/v2"
)

type Youtube struct{}

func (y *Youtube) Listen(wg *sync.WaitGroup, streamLink string, pipes pipe.Pipes, ch pipe.MsgStream) {
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
		for _, msg := range chat {
			ch <- pipe.Message{Nickname: msg.AuthorName, Text: msg.Message, Platform: pipe.YoutubePlatform}
		}
	}
}
