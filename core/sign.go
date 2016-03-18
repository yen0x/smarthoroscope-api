package core

import "strconv"

type Sign struct {
	Name string `json:"name"`

	Categories []Category `json:"categories"`

	Date string `json:"date"`
}

func (s Sign) Generate(name, date string) *Sign {
	var cat Category
	s.Categories = make([]Category, 4)
	for i, _ := range s.Categories {
		s.Categories[i] = cat.Generate()
		s.Categories[i].Id = strconv.Itoa(i)
	}
	s.Name = name
	//now := time.Now()
	//date := (time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC))
	//s.Date = date.Format(DATE_FORMAT)
	s.Date = date
	return &s
}
