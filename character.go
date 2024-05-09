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
	athletics     Attribute `pl:"Atletyka"`
	vigilance     Attribute `pl:"Czujność"`
	determination Attribute `pl:"Determinacja"`
	fortidude     Attribute `pl:"Hart"`
	intuition     Attribute `pl:"Intuicja"`
	strength      Attribute `pl:"Krzepa"`
	medics        Attribute `pl:"Medycyna"`
	taming        Attribute `pl:"Obłaskawianie"`
	religiousness Attribute `pl:"Pobożność"`
}

func newAttributes() Attributes {
	return Attributes{
		athletics:     Attribute{name: "Atletyka", proficiency: 0, focus: 0},
		vigilance:     Attribute{name: "Czujność", proficiency: 0, focus: 0},
		determination: Attribute{name: "Determinacja", proficiency: 0, focus: 0},
		fortidude:     Attribute{name: "Hart", proficiency: 0, focus: 0},
		intuition:     Attribute{name: "Intuicja", proficiency: 0, focus: 0},
		strength:      Attribute{name: "Krzepa", proficiency: 0, focus: 0},
		medics:        Attribute{name: "Medycyna", proficiency: 0, focus: 0},
		taming:        Attribute{name: "Obłaskawianie", proficiency: 0, focus: 0},
		religiousness: Attribute{name: "Pobożność", proficiency: 0, focus: 0},
	}
}

func (c Character) Name() string {
	return c.name
}
