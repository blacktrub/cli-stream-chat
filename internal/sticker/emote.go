package sticker

type Emote interface {
	name() string
	path() string
	filename() string
	IsSupported() bool
	CheckIfExists() error
}
