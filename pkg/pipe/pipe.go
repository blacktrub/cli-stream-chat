package pipe

// TODO: maybe platform must be in separate package
type Platform struct {
	Name string
}

type Message struct {
	Nickname string
	Text     string
	Platform Platform
}

type Pipe interface {
	Write(Message)
}

type Pipes []Pipe

const Twitch string = "TW"
const Youtube string = "YT"

func WriteAll(pipes Pipes, msg Message) {
	for i := 0; i < len(pipes); i++ {
		pipes[i].Write(msg)
	}
}
