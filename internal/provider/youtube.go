package provider

import (
	"cli-stream-chat/internal/msg"
	"context"
	"fmt"

	yt "github.com/abhinavxd/youtube-live-chat-downloader/v2"
)

type Youtube struct {
	stream string
}

func NewYoutubeProvider(stream string) *Youtube {
	return &Youtube{
		stream: stream,
	}
}

func (y *Youtube) Listen(ctx context.Context, out chan msg.Message) error {
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
		for _, m := range chat {
			out <- *msg.NewYoutube(m.AuthorName, m.Message)
		}
	}
}
