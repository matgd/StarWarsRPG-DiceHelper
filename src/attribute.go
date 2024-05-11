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

func (a *Attribute) SetProficiency(value int) {
	a.proficiency = value
}

func (a Attribute) Focus() int {
	return a.focus
}

func (a *Attribute) SetFocus(value int) {
	a.focus = value
}
