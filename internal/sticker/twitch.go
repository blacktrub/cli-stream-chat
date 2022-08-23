/*
Image URL
https://static-cdn.jtvnw.net/emoticons/v2/196892/static/light/1.0
*/

package sticker

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type TwitchEmote struct {
	ID   string
	Name string
}

func (e TwitchEmote) name() string {
	return e.Name
}

func (e TwitchEmote) filename() string {
	return filepath.Join(StickersPath, e.Name+".png")
}

func (e TwitchEmote) path() string {
	return fmt.Sprintf("https://static-cdn.jtvnw.net/emoticons/v2/%s/static/light/2.0", e.ID)
}

func (e TwitchEmote) IsSupported() bool {
	return true
}

func (e TwitchEmote) CheckIfExists() error {
	_, err := os.ReadFile(e.filename())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = Download(e)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

type TwitchEmotes []TwitchEmote
