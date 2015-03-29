package core

type Sign struct {
	Name string `json:"name"`

	Category1 Category
	Category2 Category
	Category3 Category
	Category4 Category

	Date string `json:"date"`
}

func (s Sign) Generate(name, date string) *Sign {
	var cat Category
	s.Category1, s.Category2, s.Category3, s.Category4 = cat.Generate(), cat.Generate(), cat.Generate(), cat.Generate()
	s.Name = name
	//now := time.Now()
	//date := (time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC))
	//s.Date = date.Format(DATE_FORMAT)
	s.Date = date
	return &s
}
