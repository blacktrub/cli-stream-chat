/*
https://gist.github.com/chuckxD/377211b3dd3e8ca8dc505500938555eb

fetch channel stickers
https://api.betterttv.net/3/cached/users/twitch/571574557

fetch global stickers
https://api.betterttv.net/3/cached/emotes/global

fetch sticker
https://cdn.betterttv.net/emote/5a970ab2122e4331029f0d7e/3x
*/
package sticker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: it's Twitch ID
var btvUserId = 571574557

type BTTVSticker struct {
	id   string
	code string
}

type channelStickersResponse struct {
	Avatar        string        `json:"avatar"`
	Bots          []interface{} `json:"bots"`
	ChannelEmotes []interface{} `json:"channelEmotes"`
	ID            string        `json:"id"`
	SharedEmotes  []struct {
		Code      string `json:"code"`
		ID        string `json:"id"`
		ImageType string `json:"imageType"`
		User      struct {
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			ProviderID  string `json:"providerId"`
		} `json:"user"`
	} `json:"sharedEmotes"`
}

type globalStickersResponse []struct {
	Code      string `json:"code"`
	ID        string `json:"id"`
	ImageType string `json:"imageType"`
	UserID    string `json:"userId"`
}

// TODO: quite the same code for fetch global and user's stickers
func getGlobalStickers() []BTTVSticker {
	resp, err := http.Get("https://api.betterttv.net/3/cached/emotes/global")
	if err != nil {
		return []BTTVSticker{}
	}
	defer resp.Body.Close()
	var data globalStickersResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return []BTTVSticker{}
	}
	var stickers []BTTVSticker
	for i := 0; i < len(data); i++ {
		s := data[i]
		stickers = append(stickers, BTTVSticker{s.ID, s.Code})
	}
	return stickers

}

func getUserStickers(userId int) []BTTVSticker {
	resp, err := http.Get(fmt.Sprintf("https://api.betterttv.net/3/cached/users/twitch/%d", userId))
	if err != nil {
		return []BTTVSticker{}
	}
	defer resp.Body.Close()
	var data channelStickersResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return []BTTVSticker{}
	}
	var stickers []BTTVSticker
	for i := 0; i < len(data.SharedEmotes); i++ {
		s := data.SharedEmotes[i]
		stickers = append(stickers, BTTVSticker{s.ID, s.Code})
	}
	return stickers
}

// TODO: cache it for a while
func GetBTTVStickers() []BTTVSticker {
	var stickers []BTTVSticker
	stickers = append(stickers, getGlobalStickers()...)
	stickers = append(stickers, getUserStickers(btvUserId)...)
	return stickers
}
