package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/ricochhet/minicommon/filesystem"
	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/info"
	"github.com/ricochhet/modmanager/manager"
	"github.com/ricochhet/modmanager/pkg/logger"
)

var (
	defopts *aflag.Options = aflag.NewOptions() //nolint:gochecknoglobals // wontfix
	opts    *aflag.Options = aflag.NewOptions() //nolint:gochecknoglobals // wontfix
)

var (
	gitHash   string //nolint:gochecknoglobals // wontfix
	buildDate string //nolint:gochecknoglobals // wontfix
	buildOn   string //nolint:gochecknoglobals // wontfix
)

func printVersion() {
	logger.SharedLogger.Info(buildOn)
	logger.SharedLogger.Infof("GitHash: %s", gitHash)
	logger.SharedLogger.Infof("Build Date: %s", buildDate)
}

func main() {
	logfile := aflag.OpenLogFile(*opts)
	defer func() {
		if err := logfile.Close(); err != nil {
			log.Fatalf("Error closing logfile: %v", err)
		}
	}()

	logger.SharedLogger = logger.NewLogger(logger.InfoLevel, io.MultiWriter(logfile, os.Stdout), log.Lshortfile|log.LstdFlags)

	versionFlag := flag.Bool("v", false, "Print the current version")

	configfile := aflag.OpenConfigFile(*opts)
	defer func() {
		if err := configfile.Close(); err != nil {
			log.Fatalf("Error closing configfile: %v", err)
		}
	}()

	configData, err := filesystem.ReadAllLines(configfile)
	if err != nil {
		logger.SharedLogger.Fatalf("Error reading config file: %v", err)
	}

	config, err := aflag.MapConfigFile(configData)
	if err != nil {
		logger.SharedLogger.Fatalf("Error reading config key value pairs: %v", err)
	}

	aflag.StrVar(&opts.Game, "game", defopts.Game, "The selected game", config)
	aflag.StrVar(&opts.LoadOrder, "loadOrder", defopts.LoadOrder, "The load order file to use", config)
	aflag.StrVar(&opts.Addons, "addons", defopts.Addons, "The addon file to use", config)
	aflag.StrVar(&opts.Renames, "renames", defopts.Renames, "The rename file to use", config)
	aflag.StrVar(&opts.Exclusions, "exclusions", defopts.Exclusions, "The exclusion file to use", config)
	aflag.StrVar(&opts.Engine, "engine", defopts.Engine, "The engine file to use", config)
	aflag.StrVar(&opts.Formats, "formats", defopts.Formats, "The formats file to use", config)
	aflag.StrVar(&opts.Bin, "bin", defopts.Bin, "The 7z binary file path", config)
	aflag.BoolVar(&opts.Silent, "silent", defopts.Silent, "Whether 7z output should be visible or not", config)
	aflag.StrVar(&opts.Data, "data", defopts.Data, "The data directory to use", config)
	aflag.StrVar(&opts.Mods, "mods", defopts.Mods, "The mods directory to use", config)
	aflag.StrVar(&opts.Temp, "temp", defopts.Temp, "The temp directory to use", config)
	aflag.StrVar(&opts.Output, "output", defopts.Output, "The output directory", config)
	aflag.StrVar(&opts.User, "user", defopts.User, "The user directory to use", config)
	aflag.StrVar(&opts.Hook, "hook", defopts.Hook, "The process hook format (.dll)", config)
	aflag.StrVar(&opts.Config, "config", defopts.Config, "The config file to read from", config)
	aflag.StrVar(&opts.Log, "log", defopts.Log, "The log file to write to", config)

	flag.Parse()

	info.DrawWatermark(info.WatermarkText())

	if *versionFlag {
		printVersion()
		return
	}

	if err := manager.Setup(*opts); err != nil {
		logger.SharedLogger.Fatalf("Error during setup: %v", err)
	}

	if err := manager.Process(*opts); err != nil {
		logger.SharedLogger.Fatalf("Error during processing: %v", err)
	}

	if err := manager.CleanTemp(*opts); err != nil {
		logger.SharedLogger.Warnf("Could not clean temp directory: %v", err)
	}

	if err := manager.CleanEmpty(*opts); err != nil {
		logger.SharedLogger.Warnf("Could not clean empty directories: %v", err)
	}
}
