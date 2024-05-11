package main

import "fmt"

type LocalePL interface {
	Save() string
	Filter3Dots() string
	DodatkoweKosci() string
	SavingSuccessful() string
	ShowChancesForSuccess() string
	AtLeastNonMdices(n int, m int) string
	ChancesDontAccountFocus() string
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

func (l Locale) SavingSuccessful() string {
	return "Zapisano pomyślnie"
}

func (l Locale) ShowChancesForSuccess() string {
	return "Pokaż szanse na sukces"
}

func (l Locale) AtLeastNonMdices(n int, m int) string {
	if m == 1 {
		return "Co najmniej " + fmt.Sprint(n) + " na " + fmt.Sprint(m) + " kości"
	}
	return "Co najmniej " + fmt.Sprint(n) + " na " + fmt.Sprint(m) + " kościach"
}

func (l Locale) ChancesDontAccountFocus() string {
	return "Szansa nie uwzględnia skupienia"
}
