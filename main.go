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

const (
	APP_NAME      string  = "Star Wars PRG Dice Calculator"
	WINDOW_SIZE_X float32 = 600
	WINDOW_SIZE_Y float32 = 300
)

func coreAttributeHBoxes(playerCharacter *Character) []fyne.CanvasObject {
	coreAttributeWidgets := [][2]fyne.CanvasObject{
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.body.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.mind.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.spirit.Name()))},
	}

	return []fyne.CanvasObject{
		container.NewHBox(coreAttributeWidgets[0][0], coreAttributeWidgets[0][1]),
		container.NewHBox(coreAttributeWidgets[1][0], coreAttributeWidgets[1][1]),
		container.NewHBox(coreAttributeWidgets[2][0], coreAttributeWidgets[2][1]),
	}
}

func calculatorCoreAttributeRadio(calculator *DiceCalculator) fyne.CanvasObject {
	radioChoices := []string{string(BODY), string(MIND), string(SPIRIT)}
	coreAttributeRadio := widget.NewRadioGroup(radioChoices, func(s string) {
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

	return coreAttributeRadio
}

func attributeHBoxes(playerCharacter *Character) []fyne.CanvasObject {
	attributeWidgets := [][2]fyne.CanvasObject{
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.athletics.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.vigilance.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.determination.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.fortidude.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.intuition.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.strength.Name()))},

		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.medics.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.taming.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.religiousness.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.cunning.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.survival.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.reflexes.Name()))},

		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.craftmanship.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.stealth.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.force.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.theology.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.ranged.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.melee.Name()))},

		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.knowledge.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.secretKnowledge.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.natureKnowledge.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.entartainment.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.intimidation.Name()))},
	}

	return []fyne.CanvasObject{
		container.NewHBox(attributeWidgets[0][0], attributeWidgets[0][1]),
	}
}

func leftVBox(playerCharacter *Character, calculator *DiceCalculator) fyne.CanvasObject {
	left := []fyne.CanvasObject{}
	left = append(left, coreAttributeHBoxes(playerCharacter)...)
	left = append(left, widget.NewSeparator())
	left = append(left, calculatorCoreAttributeRadio(calculator))
	return container.NewVBox(
		left...,
	)
}

func rightVBox(playerCharacter *Character) fyne.CanvasObject {
	right := []fyne.CanvasObject{}
	right = append(right, attributeHBoxes(playerCharacter)...)
	return container.NewVBox(
		right...,
	)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&MyTheme{})

	w := a.NewWindow(APP_NAME)
	w.Resize(fyne.NewSize(WINDOW_SIZE_X, WINDOW_SIZE_Y))

	playerCharacter := NewCharacter("Gordo")
	calculator := NewDiceCalculator(&playerCharacter)

	w.SetContent(container.NewHBox(
		container.NewVBox(leftVBox(&playerCharacter, &calculator)),
		widget.NewSeparator(),
		container.NewVBox(rightVBox(&playerCharacter)),
	))

	w.ShowAndRun()
}
