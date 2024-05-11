package main

import (
	"fmt"
	"log"
	"math"
)

type DiceCalculator struct {
	character           *Character
	chosenCoreAttribute *CoreAttribute
	chosenAttribute     *Attribute
	modifier            int
}

func NewDiceCalculator(character *Character) DiceCalculator {
	return DiceCalculator{
		character:           character,
		chosenCoreAttribute: &character.coreAttributes.body,
		chosenAttribute:     &character.attributes.athletics,
		modifier:            0,
	}
}

func (dc *DiceCalculator) SetCoreAttribute(attributeName CoreAttributeName) {
	switch attributeName {
	case BODY:
		dc.chosenCoreAttribute = &dc.character.coreAttributes.body
	case MIND:
		dc.chosenCoreAttribute = &dc.character.coreAttributes.mind
	case SPIRIT:
		dc.chosenCoreAttribute = &dc.character.coreAttributes.spirit
	}
}

func (dc *DiceCalculator) SetAttributeByPlName(attributeName string) {
	if chosenAttribute, ok := dc.character.PlAttributeNamesMap()[attributeName]; ok {
		dc.chosenAttribute = chosenAttribute
		return
	}
	log.Fatalf("Attribute with name %v not found", attributeName)
}

func (dc DiceCalculator) String() string {
	return fmt.Sprintf("Character: %v\nChosen Core Attribute: %v\nChosen Attribute: %v\nModifier: %d",
		dc.character.Name(),
		dc.chosenCoreAttribute,
		dc.chosenAttribute,
		dc.modifier)
}

func (dc DiceCalculator) Print() {
	fmt.Println("\nDice Calculator:")
	fmt.Println(dc.String())
}

func (dc DiceCalculator) dies() int {
	availableDies := dc.chosenCoreAttribute.Value()
	availableDies += dc.chosenAttribute.Proficiency()
	availableDies += dc.modifier
	return availableDies
}

// Chances returns the probability of rolling at least greaterEqual successes with available dies
func (dc DiceCalculator) Chances(successDies, greaterEqual int) float64 {
	dies := float64(dc.dies())
	ge := float64(greaterEqual)

	universe := 1.0
	failThrowChance := 1.0 - ((7.0 - ge) / 6.0)
	possibilityOf0Success := math.Pow(failThrowChance, dies)
	if successDies == 1 {
		return universe - possibilityOf0Success
	}
	return 0
}

func (dc DiceCalculator) FormattedChances(successDies, greaterEqual int) string {
	return fmt.Sprintf("%.2f%%", dc.Chances(successDies, greaterEqual)*100)
}

func (dc *DiceCalculator) DiceRollText() string {
	focus := dc.chosenAttribute.Focus()
	return fmt.Sprintf(".%dd6 + %df", dc.dies(), focus)
}
