package provider

import (
	"cli-stream-chat/internal"
	"context"
	"fmt"

	yt "github.com/abhinavxd/youtube-live-chat-downloader/v2"
)

type Youtube struct {
	stream string // youtube stream link
}

func NewYoutubeProvider(stream string) *Youtube {
	return &Youtube{
		stream: stream,
	}
}

func (y *Youtube) Listen(ctx context.Context, out chan internal.Message) error {
	continuation, cfg, err := yt.ParseInitialData(y.stream)
	if err != nil {
		return fmt.Errorf("youtube error: %w", err)
	}
	for {
		chat, newContinuation, error := yt.FetchContinuationChat(continuation, cfg)
		if error != nil {
			return fmt.Errorf("youtube error: %w", err)
		}
		continuation = newContinuation
		for _, msg := range chat {
			out <- internal.Message{
				Nickname: msg.AuthorName,
				Text:     msg.Message,
				Platform: internal.YoutubePlatform,
			}
		}
	}
}
