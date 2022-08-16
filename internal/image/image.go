/*
Kitty image protocol - https://sw.kovidgoyal.net/kitty/graphics-protocol/
*/

package image

import (
	b64 "encoding/base64"
	"fmt"
	"os"
)

const NullColumns int = 0

func stringToBase64(content []byte) string {
	return b64.StdEncoding.EncodeToString(content)
}

func Build(name, path string, columns int) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return name
	}

	var out string
	for {
		var chunk []byte
		var m string
		chunkSize := 4096

		if len(content) > chunkSize {
			chunk = content[:chunkSize]
			content = content[chunkSize:]
			m = "1"
		} else {
			chunk = content
			content = []byte{}
			m = "0"
		}

		// TODO: delete hardcode
		out = out + "\033_G"
		if columns == NullColumns {
			out = out + fmt.Sprintf("m=%s,a=T,f=100,r=1;", m)
		} else {
			out = out + fmt.Sprintf("m=%s,a=T,f=100,r=1,c=%d;", m, columns)
		}
		out = out + stringToBase64(chunk)
		out = out + "\033\\"

		if len(content) == 0 {
			break
		}
	}
	return out
}
