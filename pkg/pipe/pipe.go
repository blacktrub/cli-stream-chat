package pipe

type Platform struct {
	Name string
}

type Message struct {
	Nickname string
	Text     string
	platform Platform
}
