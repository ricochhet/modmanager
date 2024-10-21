package extras

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ricochhet/minicommon/filesystem"
)

type Patch struct {
	Find     string `json:"find"`
	Replace  string `json:"replace"`
	Position string `json:"position"`
}

type PatchTable struct {
	Bytes []Patch `json:"bytes"`
}

func WritePatchTable(fileName string, data PatchTable) error {
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

func ReadPatchTable(fileName string) (PatchTable, error) {
	var table PatchTable

	jsonData, err := filesystem.ReadFile(fileName)
	if err != nil {
		return table, fmt.Errorf("error reading file: %w", err)
	}

	if err = json.Unmarshal(jsonData, &table); err != nil {
		return table, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return table, nil
}
