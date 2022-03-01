package sticker

import (
	b64 "encoding/base64"
	"os"
	"path/filepath"
)

// TODO: do not use relative path
var StickersPath = "./pic"

func stringToBase64(content []byte) string {
	return b64.StdEncoding.EncodeToString(content)
}

func readStickerFile(name string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(StickersPath, name))
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func BuildKittyStiker(name string) string {
	content, err := readStickerFile(name)
	if err != nil {
		return name
	}

	var out string
	// TODO: delete hardcode
	out = out + "\033_G"
	out = out + "m=0,a=T,f=100;"
	out = out + stringToBase64(content)
	out = out + "\033\\"
	return out
}
