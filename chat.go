package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"cli-stream-chat/pkg/pipe"

	"github.com/abhinavxd/youtube-live-chat-downloader/v2"
	"github.com/gempir/go-twitch-irc/v3"
)

var TwitchPlatform = pipe.Platform{"TW"}
var YoutubePlatform = pipe.Platform{"YT"}

func printMessage(msg pipe.Message, colorize func(string) string) {
	fmt.Println(fmt.Sprintf("%s: %s", colorize(msg.Nickname), msg.Text))
}

func makeBlue(m string) string {
	return fmt.Sprintf("\033[1;34m%s\033[0m", m)
}

func makeRed(m string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}

func listenYoutube(wg sync.WaitGroup, streamLink string) {
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
			m := pipe.Message{msg.AuthorName, msg.Message, YoutubePlatform}
			printMessage(m, makeRed)
		}
	}
}

func listenTwitch(wg sync.WaitGroup, channelName string) {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		m := pipe.Message{message.User.DisplayName, message.Message, TwitchPlatform}
		printMessage(m, makeBlue)
	})

	client.Join(channelName)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func runListeners(twitch string, youtubeLink string) {
	var wg sync.WaitGroup

	if twitch != "" {
		wg.Add(1)
		go listenTwitch(wg, twitch)
	}

	if youtubeLink != "" {
		wg.Add(1)
		go listenYoutube(wg, youtubeLink)
	}

	wg.Wait()
}

func main() {
	twitch := flag.String("twitch", "", "Twitch channel name")
	youtubeLink := flag.String("youtube", "", "Youtube stream link")
	flag.Parse()
	if *twitch == "" && *youtubeLink == "" {
		fmt.Println("Bad run arguments")
		os.Exit(0)
	}
	runListeners(*twitch, *youtubeLink)
}
