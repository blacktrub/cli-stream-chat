package pipe

import "fmt"

type Stdout struct{}

func (s *Stdout) Write(msg Message) {
	colorize := getColorize(msg.Platform)
	fmt.Println(fmt.Sprintf("%s: %s", colorize(msg.Nickname), msg.Text))
}

func getColorize(p Platform) func(string) string {
	switch p.Name {
	case Twitch:
		return makeBlue
	case Youtube:
		return makeRed
	default:
		return withoutColor
	}
}

func makeBlue(m string) string {
	return fmt.Sprintf("\033[1;34m%s\033[0m", m)
}

func makeRed(m string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", m)
}

func withoutColor(m string) string {
	return m
}
