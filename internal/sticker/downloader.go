package sticker

import (
	"io"
	"net/http"
	"os"
)

func Download(emote Emote) error {
	resp, err := http.Get(emote.path())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(emote.filename())
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
