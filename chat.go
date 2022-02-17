package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"cli-stream-chat/pkg/pipe"

	"github.com/abhinavxd/youtube-live-chat-downloader/v2"
	"github.com/gempir/go-twitch-irc/v3"
)

var TwitchPlatform = pipe.Platform{Name: pipe.Twitch}
var YoutubePlatform = pipe.Platform{Name: pipe.Youtube}

func printMessage(msg pipe.Message, colorize func(string) string) {
	fmt.Println(fmt.Sprintf("%s: %s", colorize(msg.Nickname), msg.Text))
}

// TODO: create a module for platforms
func listenYoutube(wg sync.WaitGroup, streamLink string, pipes pipe.Pipes) {
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
			m := pipe.Message{Nickname: msg.AuthorName, Text: msg.Message, Platform: YoutubePlatform}
			pipe.WriteAll(pipes, m)
		}
	}
}

func listenTwitch(wg sync.WaitGroup, channelName string, pipes pipe.Pipes) {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		m := pipe.Message{Nickname: message.User.DisplayName, Text: message.Message, Platform: TwitchPlatform}
		pipe.WriteAll(pipes, m)
	})

	client.Join(channelName)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func runListeners(twitch string, youtubeLink string, pipes pipe.Pipes) {
	var wg sync.WaitGroup

	if twitch != "" {
		wg.Add(1)
		go listenTwitch(wg, twitch, pipes)
	}

	if youtubeLink != "" {
		wg.Add(1)
		go listenYoutube(wg, youtubeLink, pipes)
	}

	wg.Wait()
}

func main() {
	twitch := flag.String("twitch", "", "Twitch channel name")
	youtubeLink := flag.String("youtube", "", "Youtube stream link")
	keepLog := flag.Bool("keep-log", false, "Keep stream log")
	devices := flag.String("devices", "", "List tty devices")
	flag.Parse()

	if *twitch == "" && *youtubeLink == "" {
		fmt.Println("Bad run arguments")
		os.Exit(0)
	}

	pipes := pipe.Pipes{&pipe.Stdout{}}
	if *keepLog {
        // TODO: move logs file path to a flag, if it is present, then write log
		pipes = append(pipes, &pipe.Log{Path: "/home/bt/stream/log"})
	}

	if *devices != "" {
		deviceArr := strings.Split(*devices, ",")
		for i := 0; i < len(deviceArr); i++ {
			pipes = append(pipes, &pipe.Device{Path: deviceArr[i]})
		}
	}

	runListeners(*twitch, *youtubeLink, pipes)
}
