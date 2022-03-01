package pipe

import (
	"cli-stream-chat/pkg/msg"
)

type Pipe interface {
	Write(msg.Message)
}

type Pipes []Pipe

func WriteAll(pipes Pipes, msg msg.Message) {
	for i := 0; i < len(pipes); i++ {
		pipes[i].Write(msg)
	}
}
