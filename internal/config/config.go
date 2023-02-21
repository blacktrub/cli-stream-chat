package config

import (
	"errors"
	"flag"
	"strings"
)

type Config struct {
	Twitch      string
	YoutubeLink string
	LogPath     string
	Devices     []string
}

func (c *Config) Valid() error {
	// TODO: validate log path
	if c.Twitch == "" && c.YoutubeLink == "" {
		return errors.New("Setup at least one provider")
	}
	return nil
}

func New() Config {
	twitch := flag.String("twitch", "", "Twitch channel name")
	youtubeLink := flag.String("youtube", "", "Youtube stream link")
	logPath := flag.String("log", "", "Save stream log to file")
	devices := flag.String("devices", "", "List tty devices")
	flag.Parse()

	devs := []string{}
	if *devices != "" {
		devs = strings.Split(*devices, ",")
	}
	return Config{Twitch: *twitch, YoutubeLink: *youtubeLink, LogPath: *logPath, Devices: devs}
}
