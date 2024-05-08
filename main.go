package main

// Do the
// go mod tidy
// go run main.go

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CoreAttribute struct {
	name  string
	value int
}

type Attribute struct {
	name        string
	proficiency int // wyszkolenie
	focus       int // skupienie

}

type CoreAttributes struct {
	body   CoreAttribute
	mind   CoreAttribute
	spirit CoreAttribute
}

func newCoreAttributes() CoreAttributes {
	return CoreAttributes{
		body:   CoreAttribute{name: "Ciało", value: 0},
		mind:   CoreAttribute{name: "Umysł", value: 0},
		spirit: CoreAttribute{name: "Dusza", value: 0},
	}
}

type Attributes struct {
	athletics Attribute
	vigilance Attribute
}

func newAttributes() Attributes {
	return Attributes{
		athletics: Attribute{name: "Atletyka", proficiency: 0, focus: 0},
		vigilance: Attribute{name: "Czujność", proficiency: 0, focus: 0},
	}
}

type Character struct {
	name           string
	coreAttributes CoreAttributes
	attributes     Attributes
}

func newCharacter(name string) Character {
	return Character{
		name:           name,
		coreAttributes: newCoreAttributes(),
		attributes:     newAttributes(),
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(400, 300))

	character := newCharacter("Gordo")

	chosenCoreAttributeRadio := widget.NewRadioGroup([]string{"Ciało", "Umysł", "Dusza"}, func(s string) {
		switch s {
		case "Ciało":
			character.coreAttributes.body.value++
		case "Umysł":
			character.coreAttributes.mind.value++
		case "Dusza":
			character.coreAttributes.spirit.value++
		}
	})
	coreAttributeWidgets := [][3]fyne.CanvasObject{
		{widget.NewEntry(), widget.NewLabel(character.coreAttributes.body.name)},
		{widget.NewEntry(), widget.NewLabel(character.coreAttributes.mind.name)},
		{widget.NewEntry(), widget.NewLabel(character.coreAttributes.spirit.name)},
	}

	w.SetContent(container.NewVBox(
		container.NewHBox(coreAttributeWidgets[0][0], coreAttributeWidgets[0][1]),
		container.NewHBox(coreAttributeWidgets[1][0], coreAttributeWidgets[1][1]),
		container.NewHBox(coreAttributeWidgets[2][0], coreAttributeWidgets[2][1]),
		widget.NewSeparator(),
		chosenCoreAttributeRadio,
	))

	w.ShowAndRun()
}
