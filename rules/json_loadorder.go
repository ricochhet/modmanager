package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricochhet/simplefs"
)

type LoadOrder struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type JSONLoadOrder struct {
	JSON []LoadOrder `json:"loadOrder"`
}

func WriteLoadOrders(fileName string, data JSONLoadOrder) error {
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

func ReadLoadOrders(filePath string) (JSONLoadOrder, error) {
	var jsonData JSONLoadOrder

	data, err := simplefs.ReadFile(filePath)
	if err != nil {
		return jsonData, fmt.Errorf("error reading file: %w", err)
	}

	if err = json.Unmarshal(data, &jsonData); err != nil {
		return jsonData, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return jsonData, nil
}
