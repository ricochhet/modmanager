package manager

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/otiai10/copy"
	"github.com/ricochhet/minicommon/filesystem"
	"github.com/ricochhet/minicommon/util"
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/modmanager/rules"
)

var errIndexOutOfRange = errors.New("load order index is out of range")

//nolint:funlen,gocognit,gocyclo,cyclop,lll // wontfix
func generate(opt aflag.Options, loadOrder rules.JSONLoadOrder, addons rules.JSONAddons, renames rules.JSONRenames, exclusions rules.JSONExclusions) error {
	dirs, err := os.ReadDir(aflag.TempPath(opt))
	if err != nil {
		return err
	}

	if len(dirs) == 0 {
		return errNoFiles
	}

	game, err := FindGame(opt)
	if err != nil {
		return err
	}

	if err := CleanOutput(opt); err != nil {
		return err
	}

	stringDirs := []string{}

	for _, entry := range dirs {
		if entry.IsDir() {
			stringDirs = append(stringDirs, entry.Name())
		}
	}

	stringDirs, err = sortLoadOrder(loadOrder, stringDirs)
	if err != nil {
		return err
	}

	for _, dir := range stringDirs {
		path := filepath.Join(aflag.TempPath(opt), dir)
		search, item, err := Search((path), game.Engine.Paths)
		exclusions := rules.Exclude(exclusions, dir, path)

		if err != nil {
			return err
		}

		if item.Unsupported {
			continue
		}

		add, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		skip := true

		for _, file := range add {
			if file.IsDir() {
				skip = false
			}

			if !file.IsDir() && filesystem.GetFileExtension(file.Name()) == opt.Hook {
				for _, hook := range game.Engine.Hooks {
					dll := filepath.Join(path, file.Name())

					if filepath.Base(dll) == hook.Name && hook.Arch == "x64" {
						dllDest := filepath.Join(aflag.OutputPath(opt), filepath.Join(hook.Requires...), hook.Dll)

						logger.SharedLogger.Infof("Copying: '%s' to '%s' (%s)", dll, dllDest, dir)

						if err := filesystem.Copy(dll, dllDest, copy.Options{ //nolint:exhaustruct // wontfix
							Skip: func(_ os.FileInfo, src, _ string) (bool, error) {
								if slices.Contains(exclusions, src) {
									logger.SharedLogger.Infof("Skipping: %s", src)

									return true, nil
								}

								return false, nil
							},
						}); err != nil {
							return err
						}
					}
				}
			} else {
				skip = false
			}
		}

		if skip {
			continue
		}

		dest := dest(opt, item)

		logger.SharedLogger.Infof("Copying: '%s' to '%s' (%s)", search, dest, dir)

		if err := filesystem.Copy(search, dest, copy.Options{ //nolint:exhaustruct // wontfix
			RenameDestination: func(_, dest string) (string, error) {
				for _, rename := range renames.JSON {
					if rename.Name == dir {
						logger.SharedLogger.Infof("Renaming: %s, %s to %s", dest, rename.Old, rename.New)

						return strings.ReplaceAll(dest, rename.Old, rename.New), nil
					}
				}

				return dest, nil
			},
			Skip: func(_ os.FileInfo, src, _ string) (bool, error) {
				if slices.Contains(exclusions, src) {
					logger.SharedLogger.Infof("Skipping: %s", src)

					return true, nil
				}

				return false, nil
			},
		}); err != nil {
			return err
		}

		if err := copyAddons(opt, addons, renames, dir, path); err != nil {
			return err
		}
	}

	return postCopyAddons(opt, addons)
}

func dest(opt aflag.Options, item aflag.Data) string {
	if len(item.Requires) == 0 {
		return filepath.Join(aflag.OutputPath(opt), item.Path)
	}

	return filepath.Join(aflag.OutputPath(opt), filepath.Join(item.Requires...))
}

func sortLoadOrder(loadOrder rules.JSONLoadOrder, dirs []string) ([]string, error) {
	for _, order := range loadOrder.JSON {
		ind := order.Index

		if ind > len(dirs) {
			return nil, errIndexOutOfRange
		}

		if ind == -1 {
			ind = len(dirs)
		}

		dirs = util.MoveEntry(dirs, order.Name, ind)
	}

	return dirs, nil
}

func copyAddons(opt aflag.Options, addons rules.JSONAddons, renames rules.JSONRenames, dir, path string) error {
	for _, addon := range addons.JSON {
		if addon.Name == dir {
			addonSrc := filepath.Join(path, addon.Source)
			addonDest := filepath.Join(aflag.OutputPath(opt), addon.Destination)

			logger.SharedLogger.Infof("Copying: '%s' to '%s' (%s)", addonSrc, addonDest, dir)

			if err := filesystem.Copy(addonSrc, addonDest, copy.Options{ //nolint:exhaustruct // wontfix
				RenameDestination: func(_, dest string) (string, error) {
					for _, rename := range renames.JSON {
						if rename.Name == dir {
							return strings.ReplaceAll(dest, rename.Old, rename.New), nil
						}
					}

					return dest, nil
				},
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func postCopyAddons(opt aflag.Options, addons rules.JSONAddons) error {
	for _, addon := range addons.JSON {
		if addon.Name == "copy" {
			addonSrc := filepath.Join(aflag.ModPath(opt), addon.Source)
			addonDest := filepath.Join(aflag.OutputPath(opt), addon.Destination)

			logger.SharedLogger.Infof("Copying: '%s' to '%s' (copy)", addonSrc, addonDest)

			if err := filesystem.Copy(addonSrc, addonDest); err != nil {
				return err
			}
		}
	}

	return nil
}
