package extras

import (
	"errors"

	"github.com/ricochhet/minicommon/crypto"
	"github.com/ricochhet/minicommon/filesystem"
	"github.com/ricochhet/modmanager/pkg/logger"
)

var (
	errPathANoExist = errors.New("first path specified does not exist")
	errPathBNoExist = errors.New("second path specified does not exist")
)

func NewDiff(args []string) error {
	if !filesystem.Exists(args[0]) {
		return errPathANoExist
	}

	if !filesystem.Exists(args[1]) {
		return errPathBNoExist
	}

	dirA, err := crypto.HashDirectory(args[0])
	if err != nil {
		return err
	}

	dirB, err := crypto.HashDirectory(args[1])
	if err != nil {
		return err
	}

	data := crypto.DiffDirectory(dirA, dirB, args[0], args[1])

	for _, diff := range data {
		if diff.Local != (crypto.DiffLocalData{}) { //nolint:exhaustruct // wontfix
			logger.SharedLogger.Infof("File: %s exists in %s, but not in %s", diff.Local.Path, diff.Local.ExistsA, diff.Local.ExistsB)
		}

		if diff.Hashes != (crypto.DiffHashData{}) { //nolint:exhaustruct // wontfix
			logger.SharedLogger.Infof("Hashes for file: %s do not match:", diff.Hashes.File)
			logger.SharedLogger.Infof("\t%s: %s", diff.Hashes.PathA, diff.Hashes.HashA)
			logger.SharedLogger.Infof("\t%s: %s", diff.Hashes.PathB, diff.Hashes.HashB)
		}
	}

	return nil
}
