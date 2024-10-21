package internal

import (
	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/simplefs"
)

func CleanOutput(opt Options) error {
	logger.SharedLogger.Info("Cleaning output directory.")

	if err := simplefs.DeleteDirectory(OutputPath(opt)); err != nil {
		return err
	}

	return nil
}

func CleanTemp(opt Options) error {
	logger.SharedLogger.Info("Cleaning temp directory.")

	if err := simplefs.DeleteDirectory(TempPath(opt)); err != nil {
		return err
	}

	return nil
}

func CleanEmpty(opt Options) error {
	logger.SharedLogger.Info("Removing empty directories.")

	if err := simplefs.DeleteEmptyDirectories(OutputPath(opt)); err != nil {
		return err
	}

	return nil
}
