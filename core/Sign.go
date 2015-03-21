package core


type Sign struct {
	Name string

	Category1 Category
	Category2 Category
	Category3 Category
	Category4 Category
}

func (s Sign) Generate(name string) Sign {
	var cat Category
	s.Category1, s.Category2, s.Category3, s.Category4 = cat.Generate(), cat.Generate(), cat.Generate(), cat.Generate()
	s.Name = name
	return s
}
