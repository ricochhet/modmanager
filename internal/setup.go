package internal

import (
	"os"

	"github.com/ricochhet/simplefs"
)

//nolint:cyclop // wontfix
func Setup(opt Options) error {
	paths := []string{
		ModPath(opt),
		TempPath(opt),
		OutputPath(opt),
	}

	if !simplefs.Exists(LogPath(opt)) {
		if err := simplefs.WriteFile(LogPath(opt), []byte{}, os.ModePerm); err != nil {
			return err
		}
	}

	for _, path := range paths {
		if !simplefs.Exists(path) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}
	}

	if !simplefs.Exists(LoadOrderPath(opt)) {
		if err := WriteLoadOrders(LoadOrderPath(opt), JSONLoadOrder{[]LoadOrder{{"", 0}}}); err != nil {
			return err
		}
	}

	if !simplefs.Exists(AddonPath(opt)) {
		if err := WriteAddons(AddonPath(opt), JSONAddons{[]Addon{{"", "", ""}}}); err != nil {
			return err
		}
	}

	if !simplefs.Exists(RenamePath(opt)) {
		if err := WriteRenames(RenamePath(opt), JSONRenames{[]Rename{{"", "", ""}}}); err != nil {
			return err
		}
	}

	if !simplefs.Exists(ExclusionPath(opt)) {
		if err := WriteExclusions(ExclusionPath(opt), JSONExclusions{[]Exclusion{{"", ""}}}); err != nil {
			return err
		}
	}

	return nil
}
