package sticker

import (
	"cli-stream-chat/internal/detector"
	"cli-stream-chat/internal/image"
	"strings"
)

// TODO: do not use relative path
var StickersPath = "./pic/stickers"

func FindAndReplace(text string, emotes TwitchEmotes, broadcasterId string) string {
	if !detector.IsKitty() {
		return text
	}

	bttvStickers := GetBTTVStickers(broadcasterId)
	words := strings.Split(text, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]
		for _, emote := range emotes {
			if emote.Name != word {
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

			buildedSticker := buildKittySticker(emote.Name, emote.filename())
			words[i] = buildedSticker
		}

		for _, s := range bttvStickers {
			if s.Code != word {
				continue
			}
			if !s.IsSupported() {
				continue
			}

			err := s.CheckIfExists()
			if err != nil {
				// TODO: handle me
				continue
			}

			buildedSticker := buildKittySticker(s.Code, s.filename())
			words[i] = buildedSticker
		}
	}
	return strings.Join(words, " ")
}

func buildKittySticker(name, fn string) string {
	return image.Build(name, fn, image.NullColumns)

}
