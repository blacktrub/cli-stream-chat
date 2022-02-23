package provider

import (
	"cli-stream-chat/pkg/pipe"
	"sync"
)

type Provider interface {
	Listen(sync.WaitGroup, string, pipe.Pipes, pipe.MsgStream)
}
