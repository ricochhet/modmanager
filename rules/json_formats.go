package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricochhet/simplefs"
)

type JSONFormats struct {
	JSON []string `json:"formats"`
}

func WriteFormats(fileName string, data JSONFormats) error {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile(fileName, j, 0o600)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func ReadFormats(filePath string) (JSONFormats, error) {
	var jsonData JSONFormats

	data, err := simplefs.ReadFile(filePath)
	if err != nil {
		return jsonData, fmt.Errorf("error reading file: %w", err)
	}

	if err = json.Unmarshal(data, &jsonData); err != nil {
		return jsonData, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return jsonData, nil
}
