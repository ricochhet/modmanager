package manager

import (
	"github.com/ricochhet/minicommon/filesystem"
	"github.com/ricochhet/minicommon/sevenzip"
	aflag "github.com/ricochhet/modmanager/flag"
)

func extract(file string, opt aflag.Options) error {
	if filesystem.Exists(opt.Bin) {
		if _, err := sevenzip.SzBinExtract(file, aflag.TempPath(opt), opt.Bin, opt.Silent); err != nil {
			return err
		}
	} else {
		if _, err := sevenzip.SzExtract(file, aflag.TempPath(opt), opt.Silent); err != nil {
			return err
		}
	}

	return nil
}
