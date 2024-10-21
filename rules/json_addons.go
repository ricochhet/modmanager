package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricochhet/minicommon/filesystem"
)

type Addon struct {
	Name        string `json:"name"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type JSONAddons struct {
	JSON []Addon `json:"addons"`
}

func WriteAddons(fileName string, data JSONAddons) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile(fileName, jsonData, 0o600)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func ReadAddons(filePath string) (JSONAddons, error) {
	var jsonData JSONAddons

	data, err := filesystem.ReadFile(filePath)
	if err != nil {
		return jsonData, fmt.Errorf("error reading file: %w", err)
	}

	if err = json.Unmarshal(data, &jsonData); err != nil {
		return jsonData, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return jsonData, nil
}
