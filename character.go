package main

type Character struct {
	name           string
	coreAttributes CoreAttributes
	attributes     Attributes
}

func NewCharacter(name string) Character {
	return Character{
		name:           name,
		coreAttributes: newCoreAttributes(),
		attributes:     newAttributes(),
	}
}

type CoreAttributes struct {
	body   CoreAttribute
	mind   CoreAttribute
	spirit CoreAttribute
}

func newCoreAttributes() CoreAttributes {
	return CoreAttributes{
		body:   CoreAttribute{name: BODY, value: 0},
		mind:   CoreAttribute{name: MIND, value: 0},
		spirit: CoreAttribute{name: SPIRIT, value: 0},
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

func (c Character) Name() string {
	return c.name
}
