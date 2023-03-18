package color

import (
	"fmt"
)

type MakeRed struct{}

func (m MakeRed) Colorize(userId int, text string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}
