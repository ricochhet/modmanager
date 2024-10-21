package manager

import (
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/sevenzip"
	"github.com/ricochhet/simplefs"
)

func extract(file string, opt aflag.Options) error {
	if simplefs.Exists(opt.Bin) {
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
