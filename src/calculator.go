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
// You can base tests on this:
// https://www.gigacalculator.com/calculators/dice-probability-calculator.php
// At least one die >= X
func (dc DiceCalculator) Chances(successDies, greaterEqual int) float64 {
	dies := float64(dc.dies())
	ge := float64(greaterEqual)

	// At least one 6 from 3 dies
	// 1 - (5/6 * 5/6 * 5/6) = 1 - 125/216 = 91/216 ~ 42%

	// At least one 5 from 4 dies
	// 1 - (4/6 * 4/6 * 4/6 * 4/6) = 1 - 8/27 = 19/27 ~ 70%
	universe := 1.0
	failThrowChance := 1.0 - ((7.0 - ge) / 6.0)
	possibilityOf0Success := math.Pow(failThrowChance, dies)
	if successDies == 1 {
		return universe - possibilityOf0Success
	}

	// At least 2 6s from 3 dies

	// 1 minus
	// 0 6s (5/6 * 5/6 * 5/6) = 125/216

	// 1 6  (1/6 * 5/6 * 5/6) = 25/216
	// 1 6  (5/6 * 1/6 * 5/6) = 25/216
	// 1 6  (5/6 * 5/6 * 1/6) = 25/216

	// 1 - 200/216 = 16/216 ~ 7.4%
	// At least 2 6s from 4 dies
	// 1 minus
	// 0 6s (5/6 * 5/6 * 5/6 * 5/6) = 625/1296

	// 1 6 (1/6 * 5/6 * 5/6 * 5/6) = 125/1296
	// 1 6 (5/6 * 1/6 * 5/6 * 5/6) = 125/1296
	// 1 6 (5/6 * 5/6 * 1/6 * 5/6) = 125/1296
	// 1 6 (5/6 * 5/6 * 5/6 * 1/6) = 125/1296

	// 1 - 1125/1296 = 171/1296 ~ 13.2%
	successThrowChance := (7.0 - ge) / 6.0
	possibilityOf1Success := dies * (successThrowChance * math.Pow(failThrowChance, dies-1))

	if successDies == 2 {
		return universe - possibilityOf0Success - possibilityOf1Success
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
