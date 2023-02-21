package main

import (
	"context"
	"log"

	"cli-stream-chat/internal"
	"cli-stream-chat/internal/config"
	"cli-stream-chat/internal/pipe"
	"cli-stream-chat/internal/provider"
)

func main() {
	cfg := config.New()

	err := cfg.Valid()
	if err != nil {
		log.Fatalln(err.Error())
	}

	s := internal.New()
	if cfg.Twitch != "" {
		s.AddProvider(
			provider.NewTwitchProvider(cfg.Twitch),
		)
	}

	if cfg.YoutubeLink != "" {
		s.AddProvider(
			provider.NewYoutubeProvider(cfg.YoutubeLink),
		)
	}

	s.AddPipe(
		&pipe.Stdout{},
	)

	if cfg.LogPath != "" {
		s.AddPipe(
			&pipe.Log{Path: cfg.LogPath},
		)
	}

	for i := 0; i < len(cfg.Devices); i++ {
		s.AddPipe(
			&pipe.Device{Path: cfg.Devices[i]},
		)
	}

	ctx := context.Background()
	if err := s.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
