package stay_typing

import (
	"math/rand"
	"time"
)

type Tasks struct {
	Difficulty string
}

var words = []string{"aaa", "bbb", "ccc", "ddd"}

func (t Tasks) Question() string {
	rand.Seed(time.Now().UnixNano())

	return words[rand.Intn(len(words))]
}
