package gen

import (
	"math/rand"
	"time"
)

type Generator interface {
	RandomNumber(min, max int) int
}

type RandomManager struct {
}

func NewRandomManager() *RandomManager {
	return &RandomManager{}
}

func (r *RandomManager) RandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max-min) + min

	return randomNumber
}
