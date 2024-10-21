package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricochhet/minicommon/filesystem"
)

type Rename struct {
	Name string `json:"name"`
	Old  string `json:"old"`
	New  string `json:"new"`
}

type JSONRenames struct {
	JSON []Rename `json:"renames"`
}

func WriteRenames(fileName string, data JSONRenames) error {
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

func ReadRenames(filePath string) (JSONRenames, error) {
	var jsonData JSONRenames

	data, err := filesystem.ReadFile(filePath)
	if err != nil {
		return jsonData, fmt.Errorf("error reading file: %w", err)
	}

	if err = json.Unmarshal(data, &jsonData); err != nil {
		return jsonData, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return jsonData, nil
}
