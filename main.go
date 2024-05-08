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

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(400, 300))

	playerCharacter := NewCharacter("Gordo")
	calculator := NewDiceCalculator(&playerCharacter)

	chosenCoreAttributeRadio := widget.NewRadioGroup([]string{string(BODY), string(MIND), string(SPIRIT)}, func(s string) {
		switch CoreAttributeName(s) {
		case BODY:
			calculator.SetCoreAttribute(BODY)
		case MIND:
			calculator.SetCoreAttribute(MIND)
		case SPIRIT:
			calculator.SetCoreAttribute(SPIRIT)
		}
		calculator.Print()
	})
	coreAttributeWidgets := [][2]fyne.CanvasObject{
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.body.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.mind.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.spirit.Name()))},
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
