package internal

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/simplefs"
)

//nolint:cyclop // wontfix
func Search(filePath string, items []Data) (string, Data, error) {
	dirs := simplefs.GetDirectories(filePath)

	for _, dir := range dirs {
		for _, item := range items {
			fws := strings.ReplaceAll(dir, "\\", "/")
			parts := strings.Split(fws, "/")

			if item.IsDir && slices.Contains(parts, item.Path) {
				return dir, item, nil
			}
		}
	}

	files, err := os.ReadDir(filePath)
	if err != nil {
		return "", Data{}, err
	}

	for _, file := range files {
		for _, item := range items {
			path := filepath.Join(filePath, file.Name())
			ext := simplefs.GetFileExtension(path)

			if !item.IsDir && !file.IsDir() && ext == item.Path && item.Unsupported {
				logger.SharedLogger.Warnf("Unsupported file type: %s, %s", ext, path)
			}

			if !item.IsDir && !file.IsDir() && ext == item.Path {
				return filepath.Dir(path), item, nil
			}
		}
	}

	return "", Data{}, nil //nolint:exhaustruct // wontfix
}