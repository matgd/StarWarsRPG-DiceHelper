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
	APP_NAME             string  = "Star Wars PRG Dice Calculator"
	WINDOW_SIZE_X        float32 = 600
	WINDOW_SIZE_Y        float32 = 300
	ATTRIBUTE_GROUP_SIZE int     = 6
)

func coreAttributeHBoxes(playerCharacter *Character) []fyne.CanvasObject {
	widgetPairs := [][2]fyne.CanvasObject{
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.body.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.mind.Name()))},
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.coreAttributes.spirit.Name()))},
	}

	hboxes := []fyne.CanvasObject{}
	for i := 0; i < len(widgetPairs); i++ {
		hboxes = append(hboxes, container.NewHBox(widgetPairs[i][0], widgetPairs[i][1]))
	}
	return hboxes
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

func attributeVBoxes(playerCharacter *Character) []fyne.CanvasObject {
	widgets := [][2]fyne.CanvasObject{
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
		{widget.NewEntry(), widget.NewLabel(string(playerCharacter.attributes.agility.Name()))},
	}

	vboxes := []fyne.CanvasObject{}
	hboxStack := []fyne.CanvasObject{}
	for i := 0; i <= len(widgets); i++ {
		if i%ATTRIBUTE_GROUP_SIZE == 0 && i != 0 {
			vboxes = append(vboxes, container.NewVBox(hboxStack...))
			hboxStack = []fyne.CanvasObject{}
		}
		if i == len(widgets) {
			break
		}
		hboxStack = append(hboxStack, container.NewHBox(widgets[i][0], widgets[i][1]))
	}
	return vboxes
}

func searchBar() fyne.CanvasObject {
	w := widget.NewEntry()
	w.SetPlaceHolder("Filtruj...")
	// TODO: Center
	return w
}

func upperLeftVBox(playerCharacter *Character, calculator *DiceCalculator) fyne.CanvasObject {
	upperLeft := []fyne.CanvasObject{}
	upperLeft = append(upperLeft, coreAttributeHBoxes(playerCharacter)...)
	return container.NewVBox(
		upperLeft...,
	)
}

func upperRightVBox(playerCharacter *Character) fyne.CanvasObject {
	upperRight := []fyne.CanvasObject{}
	upperRight = append(upperRight, attributeVBoxes(playerCharacter)...)
	return container.NewHBox(
		upperRight...,
	)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&MyTheme{})
	pl := Locale{}

	w := a.NewWindow(APP_NAME)
	w.Resize(fyne.NewSize(WINDOW_SIZE_X, WINDOW_SIZE_Y))

	playerCharacter := NewCharacter("Gordo")
	calculator := NewDiceCalculator(&playerCharacter)

	w.SetContent(container.NewVBox(
		searchBar(),
		container.NewHBox(
			container.NewVBox(upperLeftVBox(&playerCharacter, &calculator)),
			widget.NewSeparator(),
			container.NewVBox(upperRightVBox(&playerCharacter)),
		),
		&widget.Button{Text: pl.Save()},
		widget.NewSeparator(),
		container.NewHBox(
			calculatorCoreAttributeRadio(&calculator),
			widget.NewSeparator(),
		),
	))

	w.ShowAndRun()
}
