package sticker

import (
	"cli-stream-chat/internal/detector"
	"cli-stream-chat/internal/image"
	"strings"
)

// TODO: do not use relative path
var StickersPath = "./pic/stickers"

func FindAndReplace(text string, twitchEmotes TwitchEmotes, broadcasterId string) string {
	if !detector.IsKitty() {
		return text
	}

	bttvEmotes := GetBTTVStickers(broadcasterId)
	var emotes []Emote
	// TODO: ... doesn't work
	for i := 0; i < len(twitchEmotes); i++ {
		emotes = append(emotes, twitchEmotes[i])
	}
	for i := 0; i < len(bttvEmotes); i++ {
		emotes = append(emotes, bttvEmotes[i])
	}

	words := strings.Split(text, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]
		for _, emote := range emotes {
			if emote.name() != word {
				continue
			}
			if !emote.IsSupported() {
				continue
			}
			err := emote.CheckIfExists()
			if err != nil {
				// TODO: handle me
				continue
			}

			buildedSticker := buildKittySticker(emote.name(), emote.filename())
			words[i] = buildedSticker
		}
	}
	return strings.Join(words, " ")
}

func buildKittySticker(name, fn string) string {
	return image.Build(name, fn, image.NullColumns)
}
