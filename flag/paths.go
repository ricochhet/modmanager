package flag

import (
	"path/filepath"

	"github.com/ricochhet/simplefs"
)

func ConfigPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Config)
}

func LogPath(opt Options) string {
	return simplefs.GetRelativePath(filepath.Join(opt.Data, RequiredData, opt.Log))
}

func ModPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.Mods, opt.Game)
}

func TempPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.Temp, opt.Game)
}

func OutputPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.Output, opt.Game)
}

func LoadOrderPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.LoadOrder)
}

func AddonPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.Addons)
}

func RenamePath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.Renames)
}

func ExclusionPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.Exclusions)
}
