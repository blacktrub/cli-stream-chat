package detector

import (
	"os"
)

func IsKitty() bool {
	term := os.Getenv("TERM")
	return term == "xterm-kitty"
}
