package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

type SaveData struct {
	CharacterName  string                    `json:"characterName"`
	CoreAttributes map[CoreAttributeName]int `json:"coreAttributes"`
	Attributes     map[string][2]int         `json:"attributes"`
}

const FILE_NAME string = "calculatorData.json"

func SaveToFile(calculator *DiceCalculator) error {
	file, err := os.Create(FILE_NAME)
	if err != nil {
		slog.Error("Error creating file: %v", err)
		return err
	}
	defer file.Close()

	characterName := calculator.character.Name()
	coreAttributes := map[CoreAttributeName]int{
		BODY:   calculator.character.coreAttributes.body.Value(),
		MIND:   calculator.character.coreAttributes.mind.Value(),
		SPIRIT: calculator.character.coreAttributes.spirit.Value(),
	}

	attributes := []*Attribute{
		&calculator.character.attributes.athletics,
		&calculator.character.attributes.vigilance,
		&calculator.character.attributes.determination,
		&calculator.character.attributes.fortidude,
		&calculator.character.attributes.intuition,
		&calculator.character.attributes.strength,

		&calculator.character.attributes.medics,
		&calculator.character.attributes.taming,
		&calculator.character.attributes.religiousness,
		&calculator.character.attributes.cunning,
		&calculator.character.attributes.survival,
		&calculator.character.attributes.reflexes,

		&calculator.character.attributes.craftmanship,
		&calculator.character.attributes.stealth,
		&calculator.character.attributes.force,
		&calculator.character.attributes.theology,
		&calculator.character.attributes.ranged,
		&calculator.character.attributes.melee,

		&calculator.character.attributes.knowledge,
		&calculator.character.attributes.secretKnowledge,
		&calculator.character.attributes.natureKnowledge,
		&calculator.character.attributes.entartainment,
		&calculator.character.attributes.intimidation,
		&calculator.character.attributes.agility,
	}

	attributesToSave := make(map[string][2]int)
	for _, attribute := range attributes {
		attributesToSave[attribute.Name()] = [2]int{attribute.Proficiency(), attribute.Focus()}
	}

	fileData := SaveData{
		CharacterName:  characterName,
		CoreAttributes: coreAttributes,
		Attributes:     attributesToSave,
	}

	if written, err := json.Marshal(fileData); err == nil {
		file.Write(written)
	} else {
		slog.Error("Error writing to file: %v", err)
		return err
	}

	return nil
}

func LoadFromFile() (*SaveData, error) {
	file, err := os.Open(FILE_NAME)
	if err != nil {
		slog.Error("Error opening file: %v", err)
		return &SaveData{}, err
	}
	defer file.Close()

	fileData := &SaveData{}
	if err := json.NewDecoder(file).Decode(&fileData); err != nil {
		slog.Error("Error reading from file: %v", err)
		return &SaveData{}, err
	}

	return fileData, nil
}

func (sd SaveData) RestoreCoreAttribute(name CoreAttributeName) int {
	if value, ok := sd.CoreAttributes[name]; ok {
		return value
	}
	return 0
}

func (sd SaveData) RestoreAttribute(name string) [2]int {
	if value, ok := sd.Attributes[name]; ok {
		return value
	}
	return [2]int{0, 0}
}
