package internal

import (
	"github.com/ricochhet/sevenzip"
	"github.com/ricochhet/simplefs"
)

func extract(file string, opt Options) error {
	if simplefs.Exists(opt.Bin) {
		if _, err := sevenzip.SzBinExtract(file, TempPath(opt), opt.Bin, opt.Silent); err != nil {
			return err
		}
	} else {
		if _, err := sevenzip.SzExtract(file, TempPath(opt), opt.Silent); err != nil {
			return err
		}
	}

	return nil
}
