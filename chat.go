package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/abhinavxd/youtube-live-chat-downloader/v2"
	"github.com/gempir/go-twitch-irc/v3"
)

type Platform struct {
	name string
}

type Message struct {
	nickname string
	msg      string
	platform Platform
}

var twitchPlatform = Platform{"TW"}
var youtubePlatform = Platform{"YT"}

func printMessage(msg Message, colorize func(string) string) {
	fmt.Println(fmt.Sprintf("%s: %s", colorize(msg.nickname), msg.msg))
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
			m := Message{msg.AuthorName, msg.Message, youtubePlatform}
			printMessage(m, makeRed)
		}
	}
}

func listenTwitch(wg sync.WaitGroup, channelName string) {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		m := Message{message.User.DisplayName, message.Message, twitchPlatform}
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
	wg.Add(2)

	go listenTwitch(wg, twitch)
	go listenYoutube(wg, youtubeLink)

	wg.Wait()
}

func main() {
	twitch := flag.String("twitch", "", "Twitch channel name")
	youtubeLink := flag.String("youtube", "", "Youtube stream link")
	flag.Parse()
	if *twitch == "" || *youtubeLink == "" {
		fmt.Println("Bad run arguments")
		os.Exit(0)
	}
	runListeners(*twitch, *youtubeLink)
}
