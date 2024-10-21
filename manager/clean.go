package manager

import (
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/simplefs"
)

func CleanOutput(opt aflag.Options) error {
	logger.SharedLogger.Info("Cleaning output directory.")

	if err := simplefs.DeleteDirectory(aflag.OutputPath(opt)); err != nil {
		return err
	}

	return nil
}

func CleanTemp(opt aflag.Options) error {
	logger.SharedLogger.Info("Cleaning temp directory.")

	if err := simplefs.DeleteDirectory(aflag.TempPath(opt)); err != nil {
		return err
	}

	return nil
}

func CleanEmpty(opt aflag.Options) error {
	logger.SharedLogger.Info("Removing empty directories.")

	if err := simplefs.DeleteEmptyDirectories(aflag.OutputPath(opt)); err != nil {
		return err
	}

	return nil
}
