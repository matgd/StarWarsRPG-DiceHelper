package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

type SaveData struct {
	CharacterName  string                    `json:"characterName"`
	CoreAttributes map[CoreAttributeName]int `json:"coreAttributes"`
	Attributes     map[string]int            `json:"attributes"`
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
	attributes := map[string]int{}

	fileData := SaveData{
		CharacterName:  characterName,
		CoreAttributes: coreAttributes,
		Attributes:     attributes,
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
