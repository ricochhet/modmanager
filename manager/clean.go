package manager

import (
	"github.com/ricochhet/minicommon/filesystem"
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/pkg/logger"
)

func CleanOutput(opt aflag.Options) error {
	logger.SharedLogger.Info("Cleaning output directory.")

	if err := filesystem.DeleteDirectory(aflag.OutputPath(opt)); err != nil {
		return err
	}

	return nil
}

func CleanTemp(opt aflag.Options) error {
	logger.SharedLogger.Info("Cleaning temp directory.")

	if err := filesystem.DeleteDirectory(aflag.TempPath(opt)); err != nil {
		return err
	}

	return nil
}

func CleanEmpty(opt aflag.Options) error {
	logger.SharedLogger.Info("Removing empty directories.")

	if err := filesystem.DeleteEmptyDirectories(aflag.OutputPath(opt)); err != nil {
		return err
	}

	return nil
}
