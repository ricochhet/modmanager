package rules

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ricochhet/simplefs"
)

type Exclusion struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type JSONExclusions struct {
	JSON []Exclusion `json:"exclusions"`
}

func WriteExclusions(fileName string, data JSONExclusions) error {
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

func ReadExclusions(filePath string) (JSONExclusions, error) {
	var jsonData JSONExclusions

	data, err := simplefs.ReadFile(filePath)
	if err != nil {
		return jsonData, fmt.Errorf("error reading file: %w", err)
	}

	if err = json.Unmarshal(data, &jsonData); err != nil {
		return jsonData, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return jsonData, nil
}

func Exclude(exclusions JSONExclusions, entry, path string) []string {
	files := []string{}

	for _, exclusion := range exclusions.JSON {
		if entry == exclusion.Name { //nolint:nestif // wontfix
			if simplefs.GetFileExtension(exclusion.Path) == "" { // directory
				dir := filepath.Join(path, exclusion.Path)

				if simplefs.Exists(dir) {
					files = append(files, simplefs.GetFiles(dir)...)
				}
			} else {
				file := filepath.Join(path, exclusion.Path)

				if simplefs.Exists(file) {
					files = append(files, file)
				}
			}
		}
	}

	return files
}
