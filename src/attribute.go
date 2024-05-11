package main

type Attribute struct {
	name        string
	proficiency int // wyszkolenie
	focus       int // skupienie
}

func (a Attribute) Name() string {
	return a.name
}

func (a Attribute) Proficiency() int {
	return a.proficiency
}

func (a Attribute) Focus() int {
	return a.focus
}
