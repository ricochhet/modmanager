package manager

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/ricochhet/minicommon/filesystem"
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/pkg/logger"
)

//nolint:cyclop // wontfix
func Search(filePath string, items []aflag.Data) (string, aflag.Data, error) {
	dirs := filesystem.GetDirectories(filePath)

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
		return "", aflag.Data{}, err
	}

	for _, file := range files {
		for _, item := range items {
			path := filepath.Join(filePath, file.Name())
			ext := filesystem.GetFileExtension(path)

			if !item.IsDir && !file.IsDir() && ext == item.Path && item.Unsupported {
				logger.SharedLogger.Warnf("Unsupported file type: %s, %s", ext, path)
			}

			if !item.IsDir && !file.IsDir() && ext == item.Path {
				return filepath.Dir(path), item, nil
			}
		}
	}

	return "", aflag.Data{}, nil //nolint:exhaustruct // wontfix
}
