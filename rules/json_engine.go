package rules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricochhet/minicommon/filesystem"
	aflag "github.com/ricochhet/modmanager/flag"
)

func WriteEngine(fileName string, data aflag.Engine) error {
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

func ReadEngine(filePath string) (aflag.Engine, error) {
	var jsonData aflag.Engine

	data, err := filesystem.ReadFile(filePath)
	if err != nil {
		return jsonData, fmt.Errorf("error reading file: %w", err)
	}

	if err = json.Unmarshal(data, &jsonData); err != nil {
		return jsonData, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return jsonData, nil
}
