package main

import (
	"fmt"
	"log"
	"reflect"
)

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
	athletics       Attribute `pl:"Atletyka"`
	vigilance       Attribute `pl:"Czujność"`
	determination   Attribute `pl:"Determinacja"`
	fortidude       Attribute `pl:"Hart"`
	intuition       Attribute `pl:"Intuicja"`
	strength        Attribute `pl:"Krzepa"`
	medics          Attribute `pl:"Medycyna"`
	taming          Attribute `pl:"Obłaskawianie"`
	religiousness   Attribute `pl:"Pobożność"`
	cunning         Attribute `pl:"Przebiegłość"`
	survival        Attribute `pl:"Przetrwanie"`
	reflexes        Attribute `pl:"Refleks"`
	craftmanship    Attribute `pl:"Rzemiosło"`
	stealth         Attribute `pl:"Skradanie"`
	force           Attribute `pl:"Moc"`
	theology        Attribute `pl:"Teologia"`
	ranged          Attribute `pl:"Walka dystansowa"`
	melee           Attribute `pl:"Walka wręcz"`
	knowledge       Attribute `pl:"Wiedza"`
	secretKnowledge Attribute `pl:"Wiedza tajemna"`
	natureKnowledge Attribute `pl:"Wiedza o przyrodzie"`
	entartainment   Attribute `pl:"Zabawianie"`
	intimidation    Attribute `pl:"Zastraszanie"`
	agility         Attribute `pl:"Zręczność"`
}

func getPolishName(attributeName string) string {
	f, ok := reflect.TypeOf(Attributes{}).FieldByName(attributeName)
	if !ok {
		log.Fatalf("Field %s not found in Attributes", attributeName)
	}
	return f.Tag.Get("pl")
}

func newAttributes() Attributes {
	pl := getPolishName

	return Attributes{
		athletics:       Attribute{name: pl("athletics"), proficiency: 0, focus: 0},
		vigilance:       Attribute{name: pl("vigilance"), proficiency: 0, focus: 0},
		determination:   Attribute{name: pl("determination"), proficiency: 0, focus: 0},
		fortidude:       Attribute{name: pl("fortidude"), proficiency: 0, focus: 0},
		intuition:       Attribute{name: pl("intuition"), proficiency: 0, focus: 0},
		strength:        Attribute{name: pl("strength"), proficiency: 0, focus: 0},
		medics:          Attribute{name: pl("medics"), proficiency: 0, focus: 0},
		taming:          Attribute{name: pl("taming"), proficiency: 0, focus: 0},
		religiousness:   Attribute{name: pl("religiousness"), proficiency: 0, focus: 0},
		cunning:         Attribute{name: pl("cunning"), proficiency: 0, focus: 0},
		survival:        Attribute{name: pl("survival"), proficiency: 0, focus: 0},
		reflexes:        Attribute{name: pl("reflexes"), proficiency: 0, focus: 0},
		craftmanship:    Attribute{name: pl("craftmanship"), proficiency: 0, focus: 0},
		stealth:         Attribute{name: pl("stealth"), proficiency: 0, focus: 0},
		force:           Attribute{name: pl("force"), proficiency: 0, focus: 0},
		theology:        Attribute{name: pl("theology"), proficiency: 0, focus: 0},
		ranged:          Attribute{name: pl("ranged"), proficiency: 0, focus: 0},
		melee:           Attribute{name: pl("melee"), proficiency: 0, focus: 0},
		knowledge:       Attribute{name: pl("knowledge"), proficiency: 0, focus: 0},
		secretKnowledge: Attribute{name: pl("secretKnowledge"), proficiency: 0, focus: 0},
		natureKnowledge: Attribute{name: pl("natureKnowledge"), proficiency: 0, focus: 0},
		entartainment:   Attribute{name: pl("entartainment"), proficiency: 0, focus: 0},
		intimidation:    Attribute{name: pl("intimidation"), proficiency: 0, focus: 0},
		agility:         Attribute{name: pl("agility"), proficiency: 0, focus: 0},
	}
}

func (c Character) Name() string {
	return c.name
}

func (c Character) Print() {
	fmt.Println(c.name)
	fmt.Println(c.coreAttributes)
	fmt.Println(c.attributes)
}

func (c Character) PlAttributeNamesMap() map[string]*Attribute {
	m := map[string]*Attribute{
		"Atletyka":     &c.attributes.athletics,
		"Czujność":     &c.attributes.vigilance,
		"Determinacja": &c.attributes.determination,
		"Hart":         &c.attributes.fortidude,
		"Intuicja":     &c.attributes.intuition,
		"Krzepa":       &c.attributes.strength,

		"Medycyna":      &c.attributes.medics,
		"Obłaskawianie": &c.attributes.taming,
		"Pobożność":     &c.attributes.religiousness,
		"Przebiegłość":  &c.attributes.cunning,
		"Przetrwanie":   &c.attributes.survival,
		"Refleks":       &c.attributes.reflexes,

		"Rzemiosło":        &c.attributes.craftmanship,
		"Skradanie":        &c.attributes.stealth,
		"Moc":              &c.attributes.force,
		"Teologia":         &c.attributes.theology,
		"Walka dystansowa": &c.attributes.ranged,
		"Walka wręcz":      &c.attributes.melee,

		"Wiedza":              &c.attributes.knowledge,
		"Wiedza tajemna":      &c.attributes.secretKnowledge,
		"Wiedza o przyrodzie": &c.attributes.natureKnowledge,
		"Zabawianie":          &c.attributes.entartainment,
		"Zastraszanie":        &c.attributes.intimidation,
		"Zręczność":           &c.attributes.agility,
	}
	return m
}

func PlAttributeNames() []string {
	return []string{
		"Atletyka",
		"Czujność",
		"Determinacja",
		"Hart",
		"Intuicja",
		"Krzepa",

		"Medycyna",
		"Obłaskawianie",
		"Pobożność",
		"Przebiegłość",
		"Przetrwanie",
		"Refleks",

		"Rzemiosło",
		"Skradanie",
		"Moc",
		"Teologia",
		"Walka dystansowa",
		"Walka wręcz",

		"Wiedza",
		"Wiedza tajemna",
		"Wiedza o przyrodzie",
		"Zabawianie",
		"Zastraszanie",
		"Zręczność",
	}
}
