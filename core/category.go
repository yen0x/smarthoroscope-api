package core

import(
	"math/rand"
)

type Category struct {
	Main_score int `json:"main_score"`
	Score1 int `json:"score1"`
	Score2 int `json:"score2"`
	Score3 int `json:"score3"`
	Score4 int `json:"score4"`
}

func (c Category) Generate() Category {
	c.Score1, c.Score2, c.Score3, c.Score4 = rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)
	c.Main_score = (c.Score1 + c.Score2 + c.Score3 + c.Score4) / 4
	return c
}
