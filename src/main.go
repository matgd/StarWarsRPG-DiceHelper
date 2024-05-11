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
	l := widget.NewLabel
	e := widget.NewEntry

	widgets := [][3]fyne.CanvasObject{
		{l(string(playerCharacter.attributes.athletics.Name())), e(), e()},
		{l(string(playerCharacter.attributes.vigilance.Name())), e(), e()},
		{l(string(playerCharacter.attributes.determination.Name())), e(), e()},
		{l(string(playerCharacter.attributes.fortidude.Name())), e(), e()},
		{l(string(playerCharacter.attributes.intuition.Name())), e(), e()},
		{l(string(playerCharacter.attributes.strength.Name())), e(), e()},

		{l(string(playerCharacter.attributes.medics.Name())), e(), e()},
		{l(string(playerCharacter.attributes.taming.Name())), e(), e()},
		{l(string(playerCharacter.attributes.religiousness.Name())), e(), e()},
		{l(string(playerCharacter.attributes.cunning.Name())), e(), e()},
		{l(string(playerCharacter.attributes.survival.Name())), e(), e()},
		{l(string(playerCharacter.attributes.reflexes.Name())), e(), e()},

		{l(string(playerCharacter.attributes.craftmanship.Name())), e(), e()},
		{l(string(playerCharacter.attributes.stealth.Name())), e(), e()},
		{l(string(playerCharacter.attributes.force.Name())), e(), e()},
		{l(string(playerCharacter.attributes.theology.Name())), e(), e()},
		{l(string(playerCharacter.attributes.ranged.Name())), e(), e()},
		{l(string(playerCharacter.attributes.melee.Name())), e(), e()},

		{l(string(playerCharacter.attributes.knowledge.Name())), e(), e()},
		{l(string(playerCharacter.attributes.secretKnowledge.Name())), e(), e()},
		{l(string(playerCharacter.attributes.natureKnowledge.Name())), e(), e()},
		{l(string(playerCharacter.attributes.entartainment.Name())), e(), e()},
		{l(string(playerCharacter.attributes.intimidation.Name())), e(), e()},
		{l(string(playerCharacter.attributes.agility.Name())), e(), e()},
	}

	vboxes := []fyne.CanvasObject{}

	leftVBoxStack := []fyne.CanvasObject{}
	middleVBoxStack := []fyne.CanvasObject{}
	rightVBoxStack := []fyne.CanvasObject{}

	column := 0
	for i := 0; i <= len(widgets); i++ {
		if i%ATTRIBUTE_GROUP_SIZE == 0 && i != 0 {
			leftVBox := container.NewVBox(leftVBoxStack...)
			middleVBox := container.NewVBox(middleVBoxStack...)
			rightVBox := container.NewVBox(rightVBoxStack...)

			hboxStack := []fyne.CanvasObject{leftVBox, middleVBox, rightVBox}
			vboxes = append(vboxes, container.NewVBox(container.NewHBox(hboxStack...)))

			leftVBoxStack = []fyne.CanvasObject{}
			middleVBoxStack = []fyne.CanvasObject{}
			rightVBoxStack = []fyne.CanvasObject{}
		}
		if i == len(widgets) {
			break
		}
		leftVBoxStack = append(leftVBoxStack, widgets[i][0])
		middleVBoxStack = append(middleVBoxStack, widgets[i][1])
		rightVBoxStack = append(rightVBoxStack, widgets[i][2])
		column++

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
