package sticker

type Emote interface {
	path() string
	filename() string
	IsSupported() bool
	CheckIfExists() error
}
