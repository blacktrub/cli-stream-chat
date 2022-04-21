package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"cli-stream-chat/internal"
	"cli-stream-chat/internal/pipe"
	"cli-stream-chat/internal/provider"
)

func main() {
	twitch := flag.String("twitch", "", "Twitch channel name")
	youtubeLink := flag.String("youtube", "", "Youtube stream link")
	logPath := flag.String("log", "", "Save stream log to file")
	devices := flag.String("devices", "", "List tty devices")
	flag.Parse()

	if *twitch == "" && *youtubeLink == "" {
		log.Fatalln("Setup at least one provider")
	}

	ctx := context.Background()
	s := internal.NewStream()

	if *twitch != "" {
		s.AddProvider(
			provider.NewTwitchProvider(*twitch),
		)
	}

	if *youtubeLink != "" {
		s.AddProvider(
			provider.NewYoutubeProvider(*youtubeLink),
		)
	}

	if len(s.GetProviders()) == 0 {
		log.Fatalln("Setup at least one provider")
	}

	s.AddPipe(
		&pipe.Stdout{},
	)

	if *logPath != "" {
		// TODO: validate log path
		s.AddPipe(
			&pipe.Log{Path: *logPath},
		)
	}

	if *devices != "" {
		deviceArr := strings.Split(*devices, ",")
		for i := 0; i < len(deviceArr); i++ {
			s.AddPipe(
				&pipe.Device{Path: deviceArr[i]},
			)
		}
	}

	if err := s.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
