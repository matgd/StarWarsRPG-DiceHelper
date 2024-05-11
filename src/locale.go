package main

type LocalePL interface {
	Save() string
	Filter3Dots() string
	DodatkoweKosci() string
}

type Locale struct {
}

var _ LocalePL = (*Locale)(nil) // Ensure Locale implements LocalePL

func (l Locale) Save() string {
	return "Zapisz"
}

func (l Locale) Filter3Dots() string {
	return "Filtruj..."
}

func (l Locale) DodatkoweKosci() string {
	return "Dodatkowe kości"
}
