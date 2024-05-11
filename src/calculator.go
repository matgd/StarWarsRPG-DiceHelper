package main

import "fmt"

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
func (dc DiceCalculator) String() string {
	return fmt.Sprintf("Character: %v\nChosen Core Attribute: %v\nChosen Attribute: %v\nModifier: %d",
		dc.character.Name(),
		dc.chosenCoreAttribute.Name(),
		dc.chosenAttribute.Name(),
		dc.modifier)
}

func (dc DiceCalculator) Print() {
	fmt.Println("\nDice Calculator:")
	fmt.Println(dc.String())
}
