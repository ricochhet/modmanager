package manager

import (
	"os"

	"github.com/ricochhet/minicommon/filesystem"
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/rules"
)

//nolint:cyclop,lll // wontfix
func Setup(opt aflag.Options) error {
	paths := []string{
		aflag.ModPath(opt),
		aflag.TempPath(opt),
		aflag.OutputPath(opt),
	}

	if !filesystem.Exists(aflag.LogPath(opt)) {
		if err := filesystem.WriteFile(aflag.LogPath(opt), []byte{}, os.ModePerm); err != nil {
			return err
		}
	}

	for _, path := range paths {
		if !filesystem.Exists(path) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}
	}

	if !filesystem.Exists(aflag.LoadOrderPath(opt)) {
		if err := rules.WriteLoadOrders(aflag.LoadOrderPath(opt), rules.JSONLoadOrder{JSON: []rules.LoadOrder{{Name: "", Index: 0}}}); err != nil {
			return err
		}
	}

	if !filesystem.Exists(aflag.AddonPath(opt)) {
		if err := rules.WriteAddons(aflag.AddonPath(opt), rules.JSONAddons{JSON: []rules.Addon{{Name: "", Source: "", Destination: ""}}}); err != nil {
			return err
		}
	}

	if !filesystem.Exists(aflag.RenamePath(opt)) {
		if err := rules.WriteRenames(aflag.RenamePath(opt), rules.JSONRenames{JSON: []rules.Rename{{Name: "", Old: "", New: ""}}}); err != nil {
			return err
		}
	}

	if !filesystem.Exists(aflag.ExclusionPath(opt)) {
		if err := rules.WriteExclusions(aflag.ExclusionPath(opt), rules.JSONExclusions{JSON: []rules.Exclusion{{Name: "", Path: ""}}}); err != nil {
			return err
		}
	}

	return nil
}
