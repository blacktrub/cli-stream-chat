package provider

import (
	"cli-stream-chat/internal/msg"
	"context"
)

type provider interface {
	Listen(context.Context, chan msg.Message) error
}

type pipe interface {
	Write(msg.Message) error
}

type Stream struct {
	providers []provider
	pipes     []pipe
}

func New() *Stream {
	return &Stream{}
}
func (s *Stream) AddProvider(providers ...provider) {
	for _, p := range providers {
		s.providers = append(s.providers, p)
	}
}

func (s *Stream) GetProviders() []provider {
	return s.providers
}

func (s *Stream) AddPipe(pipes ...pipe) {
	for _, p := range pipes {
		s.pipes = append(s.pipes, p)
	}
}

func (s *Stream) Run(ctx context.Context) error {
	out := make(chan msg.Message)
	errChan := make(chan error)
	for _, p := range s.providers {
		go func(p provider) {
			if err := p.Listen(ctx, out); err != nil {
				errChan <- err
			}
		}(p)
	}

	go func() {
		for {
			select {
			case m := <-out:
				for _, p := range s.pipes {
					go func(p pipe) {
						if err := p.Write(m); err != nil {
							errChan <- err
						}
					}(p)
				}
			}
		}
	}()

	return <-errChan
}
