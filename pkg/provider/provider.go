package provider

import (
	"cli-stream-chat/pkg/msg"
	"cli-stream-chat/pkg/pipe"
	"sync"
)

type Provider interface {
	Listen(sync.WaitGroup, string, pipe.Pipes, msg.MsgStream)
}
