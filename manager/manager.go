package manager

import (
	"errors"
	"slices"

	"github.com/ricochhet/minicommon/filesystem"
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/rules"
)

var errNoFiles = errors.New("no compatible file types in directory")

func Process(opt aflag.Options) error { //nolint:cyclop // wontfix
	files := filesystem.GetFiles(aflag.ModPath(opt))

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
		if slices.Contains(formats, filesystem.GetFileExtension(file)) {
			if err := extract(file, opt); err != nil {
				return err
			}
		}
	}

	loadOrder, err := rules.ReadLoadOrders(aflag.LoadOrderPath(opt))
	if err != nil {
		return err
	}

	addons, err := rules.ReadAddons(aflag.AddonPath(opt))
	if err != nil {
		return err
	}

	renames, err := rules.ReadRenames(aflag.RenamePath(opt))
	if err != nil {
		return err
	}

	exclusions, err := rules.ReadExclusions(aflag.ExclusionPath(opt))
	if err != nil {
		return err
	}

	return generate(opt, loadOrder, addons, renames, exclusions)
}
