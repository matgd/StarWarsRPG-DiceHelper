package main

// Do the
// go mod tidy
// go run main.go

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
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

var diceRollLabel *canvas.Text = nil // Cannot be initialized here before fyne app works
var loadedSaveData *SaveData = nil
var allTextToSearchThrough []*canvas.Text = []*canvas.Text{}

func coreAttributeHBoxes(calculator *DiceCalculator) []fyne.CanvasObject {
	e := widget.NewEntry
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
			if s == "" || s == "-" {
				attr.SetValue(0)
			} else if intValue, err := strconv.Atoi(s); err == nil {
				attr.SetValue(intValue)
			} else {
				ne.SetText("0")
			}
			diceRollLabel.Text = calculator.DiceRollText()
			diceRollLabel.Refresh()
		}

		// Load from file
		ne.SetText(fmt.Sprint(loadedSaveData.RestoreCoreAttribute(attribute.Name())))

		text := canvas.NewText(string(attribute.Name()), color.White)
		allTextToSearchThrough = append(allTextToSearchThrough, text)

		widgetPairs = append(widgetPairs, [2]fyne.CanvasObject{ne, text})
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
		diceRollLabel.Text = calculator.DiceRollText()
		diceRollLabel.Refresh()
		calculator.Print()
	}

	return radio
}

func calculatorAttributeRadio(calculator *DiceCalculator) fyne.CanvasObject {
	radioChoices := PlAttributeNames()

	// I know, I know
	// This should be probably specific layout instead of this hacky way
	// But I'm doing this for fun and I don't care about maintanence that much
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
		diceRollLabel.Text = calculator.DiceRollText()
		diceRollLabel.Refresh()
	}

	radios := []fyne.CanvasObject{}
	for _, radio := range radioPointers {
		radio.OnChanged = selectFunc
		radios = append(radios, radio)
	}

	return container.NewHBox(radios...)

}

func attributeVBoxes(calculator *DiceCalculator) fyne.CanvasObject {
	t := canvas.NewText
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

	triplets := []fyne.CanvasObject{}
	for _, attribute := range attributes {
		attr := attribute // Pointer magic boooo (seriously, good that I've read about it)
		// ^ Without this setting will only work on last element

		ne0 := e()
		ne1 := e()
		ne0.OnChanged = func(s string) {
			if s == "" || s == "-" {
				attr.SetProficiency(0)
			} else if intValue, err := strconv.Atoi(s); err == nil {
				attr.SetProficiency(intValue)
			} else {
				ne0.SetText("0")
			}
			diceRollLabel.Text = calculator.DiceRollText()
			diceRollLabel.Refresh()
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
			diceRollLabel.Text = calculator.DiceRollText()
			diceRollLabel.Refresh()
		}
		// Load from file
		restoredFromFile := loadedSaveData.RestoreAttribute(attribute.Name())
		ne1.SetText(fmt.Sprint(restoredFromFile[0]))
		ne0.SetText(fmt.Sprint(restoredFromFile[1]))

		text := t(string(attribute.Name()), color.White)
		allTextToSearchThrough = append(allTextToSearchThrough, text)

		triplets = append(triplets, container.NewHBox(text, layout.NewSpacer(), ne0, ne1))

	}
	cont := container.NewGridWithRows(6, triplets...)
	return cont
}

func searchBar() *widget.Entry {
	w := widget.NewEntry()
	w.SetPlaceHolder(pl.Filter3Dots())
	return w
}

func coreAttributeVBox(playerCharacter *Character, calculator *DiceCalculator) fyne.CanvasObject {
	coreAttrVBox := []fyne.CanvasObject{}
	coreAttrVBox = append(coreAttrVBox, coreAttributeHBoxes(calculator)...)
	return container.NewVBox(
		coreAttrVBox...,
	)
}

func allAttributesVBox(calculator *DiceCalculator) fyne.CanvasObject {
	allAttributeVBox := []fyne.CanvasObject{}
	allAttributeVBox = append(allAttributeVBox, attributeVBoxes(calculator))
	return container.NewHBox(
		allAttributeVBox...,
	)
}

func modifierVBox(calculator *DiceCalculator) fyne.CanvasObject {
	ne := widget.NewEntry()
	ne.OnChanged = func(s string) {
		if s == "" || s == "-" {
			calculator.modifier = 0
		} else if intValue, err := strconv.Atoi(s); err == nil {
			calculator.modifier = intValue
		} else {
			ne.SetText("0")
		}
		diceRollLabel.Text = calculator.DiceRollText()
		diceRollLabel.Refresh()
	}

	return container.NewVBox(
		widget.NewLabel(pl.DodatkoweKosci()),
		ne,
	)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&MyTheme{})
	diceRollLabel = canvas.NewText("", color.White)
	loadedSaveData, _ = LoadFromFile()

	w := a.NewWindow(fmt.Sprintf("%s (ver. %s)", APP_NAME, APP_VERSION))
	w.Resize(fyne.NewSize(WINDOW_SIZE_X, WINDOW_SIZE_Y))

	playerCharacter := NewCharacter("Gordo")
	calculator := NewDiceCalculator(&playerCharacter)
	diceRollLabel.TextSize = 24
	diceRollLabel.Alignment = fyne.TextAlignCenter
	diceRollLabel.Text = calculator.DiceRollText()

	saveButton := widget.NewButton(pl.Save(), func() {
		if err := SaveToFile(&calculator); err == nil {
			dialog.ShowInformation(pl.Save(), pl.SavingSuccessful(), w)
		} else {
			dialog.ShowError(err, w)
		}
	})
	search := searchBar()

	w.SetContent(container.NewVBox(
		search,
		container.NewHBox(
			coreAttributeVBox(&playerCharacter, &calculator),
			widget.NewSeparator(),
			container.NewVBox(allAttributesVBox(&calculator)),
		),
		saveButton,
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

	search.OnChanged = func(s string) {
		HighlightSearched(s, allTextToSearchThrough...)
	}

	w.ShowAndRun()
}
