package core

import (
	"math/rand"
)

type Category struct {
	Id     string `json:"id"`
	Scores []int  `json:"scores"`
}

func (c Category) Generate() Category {
	c.Scores = make([]int, 4)
	for i, _ := range c.Scores {
		c.Scores[i] = c.generateNumber()
	}
	return c
}

func (c Category) generateNumber() int {
	high, low := 10, 3
	return rand.Intn(high-low+1) + low
}
