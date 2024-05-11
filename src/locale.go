package main

type LocalePL interface {
	Save() string
}

type Locale struct {
}

var _ LocalePL = (*Locale)(nil) // Ensure Locale implements LocalePL

func (l Locale) Save() string {
	return "Zapisz"
}
