package main

import (
	"context"
	"log"

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

	s := provider.New()
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
		// TODO: err says: handle me
		file, _ := pipe.GetFile(cfg.LogPath)
		defer file.Close()

		log := pipe.NewLog(file)
		s.AddPipe(log)
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
