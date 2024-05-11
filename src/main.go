package main

// Do the
// go mod tidy
// go run main.go

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	APP_NAME             string  = "Star Wars PRG Dice Calculator"
	APP_VERSION          string  = "24.5.11"
	WINDOW_SIZE_X        float32 = 600
	WINDOW_SIZE_Y        float32 = 300
	ATTRIBUTE_GROUP_SIZE int     = 6
)

var pl LocalePL = Locale{}

var diceRollLabel = widget.NewLabel("")

func coreAttributeHBoxes(calculator *DiceCalculator) []fyne.CanvasObject {
	e := widget.NewEntry
	l := widget.NewLabel
	playerCharacter := calculator.character

	attributes := []*CoreAttribute{
		&playerCharacter.coreAttributes.body,
		&playerCharacter.coreAttributes.mind,
		&playerCharacter.coreAttributes.spirit,
	}

	hboxes := []fyne.CanvasObject{}
	widgetPairs := [][2]fyne.CanvasObject{}
	for i, attribute := range attributes {
		attr := attribute // Pointer magic boooo (seriously, good that I've read about it)
		// ^ Without this setting will only work on last element

		ne := e()
		ne.OnChanged = func(s string) {
			// TODO: Load from file
			if s == "" || s == "-" {
				attr.SetValue(0)
			} else if intValue, err := strconv.Atoi(s); err == nil {
				attr.SetValue(intValue)
			} else {
				ne.SetText("0")
			}
			diceRollLabel.SetText(calculator.DiceRollText())
		}
		widgetPairs = append(widgetPairs, [2]fyne.CanvasObject{ne, l(string(attribute.Name()))})
		hboxes = append(hboxes, container.NewHBox(widgetPairs[i][0], widgetPairs[i][1]))
	}

	return hboxes
}

func calculatorCoreAttributeRadio(calculator *DiceCalculator) fyne.CanvasObject {
	radioChoices := []string{string(BODY), string(MIND), string(SPIRIT)}
	radio := widget.NewRadioGroup(radioChoices, func(s string) {})
	var selected CoreAttributeName = BODY

	// Default choice
	radio.SetSelected(string(selected))

	radio.OnChanged = func(s string) {
		if s == "" {
			// Prevent from unselecting
			radio.SetSelected(string(selected))
			return
		}
		switch CoreAttributeName(s) {
		case BODY:
			selected = BODY
		case MIND:
			selected = MIND
		case SPIRIT:
			selected = SPIRIT
		}
		calculator.SetCoreAttribute(selected)
		diceRollLabel.SetText(calculator.DiceRollText())
		calculator.Print()
	}

	return radio
}

func calculatorAttributeRadio(calculator *DiceCalculator) fyne.CanvasObject {
	radioChoices := PlAttributeNames()
	radioPointers := []*widget.RadioGroup{
		widget.NewRadioGroup(radioChoices[:6], func(s string) {}),
		widget.NewRadioGroup(radioChoices[6:12], func(s string) {}),
		widget.NewRadioGroup(radioChoices[12:18], func(s string) {}),
		widget.NewRadioGroup(radioChoices[18:], func(s string) {}),
	}
	// Default choice
	selected := radioChoices[0]
	radioPointers[0].SetSelected(selected)

	selectFunc := func(s string) {
		if s != "" {
			// Prevent from passing empty string to calculator set
			selected = s
		}
		calculator.SetAttributeByPlName(selected)
		calculator.Print()

		for _, radio := range radioPointers {
			// Prevents from unselecting
			// Will only select if there is present in given radio
			radio.SetSelected(selected)
		}
	}

	radios := []fyne.CanvasObject{}
	for _, radio := range radioPointers {
		radio.OnChanged = selectFunc
		radios = append(radios, radio)
	}

	return container.NewHBox(radios...)

}

func attributeVBoxes(calculator *DiceCalculator) []fyne.CanvasObject {
	l := widget.NewLabel
	e := widget.NewEntry
	playerCharacter := calculator.character

	attributes := []*Attribute{
		&playerCharacter.attributes.athletics,
		&playerCharacter.attributes.vigilance,
		&playerCharacter.attributes.determination,
		&playerCharacter.attributes.fortidude,
		&playerCharacter.attributes.intuition,
		&playerCharacter.attributes.strength,

		&playerCharacter.attributes.medics,
		&playerCharacter.attributes.taming,
		&playerCharacter.attributes.religiousness,
		&playerCharacter.attributes.cunning,
		&playerCharacter.attributes.survival,
		&playerCharacter.attributes.reflexes,

		&playerCharacter.attributes.craftmanship,
		&playerCharacter.attributes.stealth,
		&playerCharacter.attributes.force,
		&playerCharacter.attributes.theology,
		&playerCharacter.attributes.ranged,
		&playerCharacter.attributes.melee,

		&playerCharacter.attributes.knowledge,
		&playerCharacter.attributes.secretKnowledge,
		&playerCharacter.attributes.natureKnowledge,
		&playerCharacter.attributes.entartainment,
		&playerCharacter.attributes.intimidation,
		&playerCharacter.attributes.agility,
	}

	vboxes := []fyne.CanvasObject{}

	leftVBoxStack := []fyne.CanvasObject{}
	middleVBoxStack := []fyne.CanvasObject{}
	rightVBoxStack := []fyne.CanvasObject{}

	widgets := [][3]fyne.CanvasObject{}
	for _, attribute := range attributes {
		attr := attribute // Pointer magic boooo (seriously, good that I've read about it)
		// ^ Without this setting will only work on last element

		ne0 := e()
		ne1 := e()
		ne0.OnChanged = func(s string) {
			// TODO: Load from file
			if s == "" || s == "-" {
				attr.SetProficiency(0)
			} else if intValue, err := strconv.Atoi(s); err == nil {
				attr.SetProficiency(intValue)
			} else {
				ne0.SetText("0")
			}
			diceRollLabel.SetText(calculator.DiceRollText())
		}
		ne1.OnChanged = func(s string) {
			// TODO: Load from file
			if s == "" || s == "-" {
				attr.SetFocus(0)
			} else if intValue, err := strconv.Atoi(s); err == nil {
				attr.SetFocus(intValue)
			} else {
				ne1.SetText("0")
			}
			diceRollLabel.SetText(calculator.DiceRollText())
		}

		widgets = append(widgets, [3]fyne.CanvasObject{l(string(attribute.Name())), ne0, ne1})
	}

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
	w.SetPlaceHolder(pl.Filter3Dots())
	// TODO: Center
	return w
}

func upperLeftVBox(playerCharacter *Character, calculator *DiceCalculator) fyne.CanvasObject {
	upperLeft := []fyne.CanvasObject{}
	upperLeft = append(upperLeft, coreAttributeHBoxes(calculator)...)
	return container.NewVBox(
		upperLeft...,
	)
}

func upperRightVBox(calculator *DiceCalculator) fyne.CanvasObject {
	upperRight := []fyne.CanvasObject{}
	upperRight = append(upperRight, attributeVBoxes(calculator)...)
	return container.NewHBox(
		upperRight...,
	)
}

func modifierVBox(calculator *DiceCalculator) fyne.CanvasObject {
	ne := widget.NewEntry()
	ne.OnChanged = func(s string) {
		if s == "" || s == "-" {
			calculator.modifier = 0
		} else if intValue, err := strconv.Atoi(s); err == nil {
			calculator.modifier = intValue
			return
		} else {
			ne.SetText("0")
		}
	}

	return container.NewVBox(
		widget.NewLabel(pl.DodatkoweKosci()),
		ne,
	)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&MyTheme{})

	w := a.NewWindow(fmt.Sprintf("%s (ver. %s)", APP_NAME, APP_VERSION))
	w.Resize(fyne.NewSize(WINDOW_SIZE_X, WINDOW_SIZE_Y))

	playerCharacter := NewCharacter("Gordo")
	calculator := NewDiceCalculator(&playerCharacter)

	w.SetContent(container.NewVBox(
		searchBar(),
		container.NewHBox(
			container.NewVBox(upperLeftVBox(&playerCharacter, &calculator)),
			widget.NewSeparator(),
			container.NewVBox(upperRightVBox(&calculator)),
		),
		&widget.Button{Text: pl.Save()},
		widget.NewSeparator(),
		container.NewHBox(
			calculatorCoreAttributeRadio(&calculator),
			widget.NewSeparator(),
			calculatorAttributeRadio(&calculator),
			widget.NewSeparator(),
			modifierVBox(&calculator),
		),
		widget.NewSeparator(),
		diceRollLabel,
	))

	w.ShowAndRun()
}
