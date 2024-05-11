package main

type CoreAttributeName string

const (
	BODY   CoreAttributeName = "Ciało"
	MIND   CoreAttributeName = "Umysł"
	SPIRIT CoreAttributeName = "Dusza"
)

type CoreAttribute struct {
	name  CoreAttributeName
	value int
}

func (ca CoreAttribute) Name() CoreAttributeName {
	return ca.name
}

func (ca CoreAttribute) Value() int {
	return ca.value
}

func (ca *CoreAttribute) SetValue(value int) {
	ca.value = value
}
