package color

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

type MakePretty struct {
	mem map[int]int
	mu  sync.Mutex
}

func NewPretty() *MakePretty {
	return &MakePretty{mem: map[int]int{}}
}

func (c *MakePretty) Colorize(userId int, nickname string) string {
	userColor, err := c.getColor(userId)
	if err != nil {
		userColor = getRandomColor()
		c.setUserColor(userId, userColor)
	}
	return fmt.Sprintf("\033[1;%dm%s\033[0m", userColor, nickname)
}

func (c *MakePretty) setUserColor(userId int, color int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mem[userId] = color
}

func (c *MakePretty) getColor(userId int) (int, error) {
	for k, v := range c.mem {
		if k == userId {
			return v, nil
		}
	}
	return 0, errors.New("empty color")
}

func getRandomColor() int {
	colors := getColors()
	i := rand.Intn(len(colors))
	return colors[i]
}

func getColors() []int {
	var colors []int
	for i := 30; i < 38; i++ {
		colors = append(colors, i)
	}
	return colors
}
