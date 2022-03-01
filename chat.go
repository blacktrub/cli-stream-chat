package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"cli-stream-chat/pkg/msg"
	"cli-stream-chat/pkg/pipe"
	"cli-stream-chat/pkg/provider"
)

func runListeners(twitch string, youtubeLink string, pipes pipe.Pipes, ch msg.MsgStream) {
	var wg sync.WaitGroup

	if twitch != "" {
		wg.Add(1)
		tw := provider.Twitch{}
		go tw.Listen(&wg, twitch, pipes, ch)
	}

	if youtubeLink != "" {
		wg.Add(1)
		yt := provider.Youtube{}
		go yt.Listen(&wg, youtubeLink, pipes, ch)
	}

	wg.Wait()
	close(ch)
}

func listenStream(ch msg.MsgStream, pipes pipe.Pipes) {
	for msg := range ch {
		pipe.WriteAll(pipes, msg)
	}
}

func main() {
	twitch := flag.String("twitch", "", "Twitch channel name")
	youtubeLink := flag.String("youtube", "", "Youtube stream link")
	logPath := flag.String("log", "", "Save stream log to file")
	devices := flag.String("devices", "", "List tty devices")
	flag.Parse()

	if *twitch == "" && *youtubeLink == "" {
		fmt.Println("Bad run arguments")
		os.Exit(0)
	}

	pipes := pipe.Pipes{&pipe.Stdout{}}
	if *logPath != "" {
		// TODO: validate log path
		pipes = append(pipes, &pipe.Log{Path: *logPath})
	}

	if *devices != "" {
		deviceArr := strings.Split(*devices, ",")
		for i := 0; i < len(deviceArr); i++ {
			pipes = append(pipes, &pipe.Device{Path: deviceArr[i]})
		}
	}

	msgStream := make(chan msg.Message)
	go listenStream(msgStream, pipes)
	runListeners(*twitch, *youtubeLink, pipes, msgStream)
}
