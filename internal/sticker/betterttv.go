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
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

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

type StickersCache struct {
	stickers []BTTVEmote
}

var cache = StickersCache{[]BTTVEmote{}}

type BTTVEmote struct {
	id   string
	Code string
	Ext  string
}

func (s BTTVEmote) filename() string {
	return filepath.Join(StickersPath, s.Code+"."+s.Ext)
}

func (s BTTVEmote) path() string {
	return fmt.Sprintf("https://cdn.betterttv.net/emote/%s/2x", s.id)
}

func (s BTTVEmote) IsSupported() bool {
	supported := [1]string{"png"}
	for _, ext := range supported {
		if ext == s.Ext {
			return true
		}
	}
	return false
}

func (s BTTVEmote) CheckIfExists() error {
	_, err := os.ReadFile(s.filename())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = Download(s)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

// TODO: quite the same code for fetch global and user's stickers
func getGlobalStickers() []BTTVEmote {
	resp, err := http.Get("https://api.betterttv.net/3/cached/emotes/global")
	if err != nil {
		return []BTTVEmote{}
	}
	defer resp.Body.Close()
	var data globalStickersResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return []BTTVEmote{}
	}
	var stickers []BTTVEmote
	for i := 0; i < len(data); i++ {
		s := data[i]
		stickers = append(stickers, BTTVEmote{s.ID, s.Code, s.ImageType})
	}
	return stickers

}

func getUserStickers(userId string) []BTTVEmote {
	resp, err := http.Get(fmt.Sprintf("https://api.betterttv.net/3/cached/users/twitch/%s", userId))
	if err != nil {
		return []BTTVEmote{}
	}
	defer resp.Body.Close()
	var data channelStickersResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return []BTTVEmote{}
	}
	var stickers []BTTVEmote
	for i := 0; i < len(data.SharedEmotes); i++ {
		s := data.SharedEmotes[i]
		stickers = append(stickers, BTTVEmote{s.ID, s.Code, s.ImageType})
	}
	return stickers
}

func GetBTTVStickers(broadcasterId string) []BTTVEmote {
	if len(cache.stickers) > 0 {
		return cache.stickers
	}

	var stickers []BTTVEmote
	stickers = append(stickers, getGlobalStickers()...)
	stickers = append(stickers, getUserStickers(broadcasterId)...)
	cache.stickers = stickers
	return stickers
}
