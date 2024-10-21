package internal

import (
	"errors"
	"slices"

	"github.com/ricochhet/simplefs"
)

var errNoFiles = errors.New("no compatible file types in directory")

func Process(opt Options) error { //nolint:cyclop // wontfix
	files := simplefs.GetFiles(ModPath(opt))

	if len(files) == 0 {
		return errNoFiles
	}

	if err := CleanTemp(opt); err != nil {
		return err
	}

	formats, err := FindFormats(opt)
	if err != nil {
		return err
	}

	for _, file := range files {
		if slices.Contains(formats, simplefs.GetFileExtension(file)) {
			if err := extract(file, opt); err != nil {
				return err
			}
		}
	}

	loadOrder, err := ReadLoadOrders(LoadOrderPath(opt))
	if err != nil {
		return err
	}

	addons, err := ReadAddons(AddonPath(opt))
	if err != nil {
		return err
	}

	renames, err := ReadRenames(RenamePath(opt))
	if err != nil {
		return err
	}

	exclusions, err := ReadExclusions(ExclusionPath(opt))
	if err != nil {
		return err
	}

	return generate(opt, loadOrder, addons, renames, exclusions)
}
