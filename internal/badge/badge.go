/*
TODO: add all of them
All badges - https://badges.twitch.tv/v1/badges/global/display
*/
package badge

import (
	"cli-stream-chat/internal/image"
	"path/filepath"
)

type Badge string

const (
	Broadcaster Badge = "broadcaster"
	Moderator         = "moderator"
	Subscriber        = "subscriber"
)

var BadgePath = "./pic/badges"

func Show(badges map[string]int) string {
	supported := getSupported()
	var out string
	for _, name := range supported {
		if _, ok := badges[string(name)]; ok {
			out = out + buildBadge(string(name))
		}
	}
	return out
}

func buildBadge(name string) string {
	path := filepath.Join(BadgePath, name+".png")
	return image.Build(name, path, 2)
}

func getSupported() []Badge {
	return []Badge{
		Broadcaster,
		Moderator,
		Subscriber,
	}
}
