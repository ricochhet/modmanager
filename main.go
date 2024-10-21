package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/ricochhet/modmanager/internal"
	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/simplefs"
)

var (
	defopts *internal.Options = internal.NewOptions() //nolint:gochecknoglobals // wontfix
	opts    *internal.Options = internal.NewOptions() //nolint:gochecknoglobals // wontfix
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

func logfile() *os.File {
	file, err := os.OpenFile(internal.LogPath(*opts), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func configfile() *os.File {
	file, err := os.OpenFile(internal.ConfigPath(*opts), os.O_CREATE|os.O_RDONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func main() { //nolint:funlen // flag setup
	logfile := logfile()
	defer func() {
		if err := logfile.Close(); err != nil {
			log.Fatalf("Error closing logfile: %v", err)
		}
	}()

	logger.SharedLogger = logger.NewLogger(logger.InfoLevel, io.MultiWriter(logfile, os.Stdout), log.Lshortfile|log.LstdFlags)

	versionFlag := flag.Bool("v", false, "Print the current version")

	flag.StringVar(&opts.Game, "game", defopts.Game, "The selected game")
	flag.StringVar(&opts.LoadOrder, "loadOrder", defopts.LoadOrder, "The load order file to use")
	flag.StringVar(&opts.Addons, "addons", defopts.Addons, "The addon file to use")
	flag.StringVar(&opts.Renames, "renames", defopts.Renames, "The rename file to use")
	flag.StringVar(&opts.Exclusions, "exclusions", defopts.Exclusions, "The exclusion file to use")
	flag.StringVar(&opts.Engine, "engine", defopts.Engine, "The engine file to use")
	flag.StringVar(&opts.Formats, "formats", defopts.Formats, "The formats file to use")
	flag.StringVar(&opts.Bin, "bin", defopts.Bin, "The 7z binary file path")
	flag.BoolVar(&opts.Silent, "silent", defopts.Silent, "Whether 7z output should be visible or not")
	flag.StringVar(&opts.Data, "data", defopts.Data, "The data directory to use")
	flag.StringVar(&opts.Mods, "mods", defopts.Mods, "The mods directory to use")
	flag.StringVar(&opts.Temp, "temp", defopts.Temp, "The temp directory to use")
	flag.StringVar(&opts.Output, "output", defopts.Output, "The output directory")
	flag.StringVar(&opts.User, "user", defopts.User, "The user directory to use")
	flag.StringVar(&opts.Hook, "hook", defopts.Hook, "The process hook format (.dll)")
	flag.StringVar(&opts.Config, "config", defopts.Config, "The config file to read from")
	flag.StringVar(&opts.Log, "log", defopts.Log, "The log file to write to")

	flag.Parse()

	configfile := configfile()
	defer func() {
		if err := configfile.Close(); err != nil {
			log.Fatalf("Error closing configfile: %v", err)
		}
	}()

	configData, err := simplefs.ReadAllLines(configfile)
	if err != nil {
		logger.SharedLogger.Fatalf("Error reading config file: %v", err)
	}

	config, err := internal.MapConfig(configData)
	if err != nil {
		logger.SharedLogger.Fatalf("Error reading config key value pairs: %v", err)
	}

	internal.SetFlagString(&opts.Game, "game", config)
	internal.SetFlagString(&opts.LoadOrder, "loadOrder", config)
	internal.SetFlagString(&opts.Addons, "addons", config)
	internal.SetFlagString(&opts.Renames, "renames", config)
	internal.SetFlagString(&opts.Exclusions, "exclusions", config)
	internal.SetFlagString(&opts.Engine, "engine", config)
	internal.SetFlagString(&opts.Formats, "formats", config)
	internal.SetFlagString(&opts.Bin, "bin", config)
	internal.SetFlagBool(&opts.Silent, "silent", config)
	internal.SetFlagString(&opts.Data, "data", config)
	internal.SetFlagString(&opts.Mods, "mods", config)
	internal.SetFlagString(&opts.Temp, "temp", config)
	internal.SetFlagString(&opts.Output, "output", config)
	internal.SetFlagString(&opts.User, "user", config)
	internal.SetFlagString(&opts.Hook, "hook", config)
	internal.SetFlagString(&opts.Config, "config", config)
	internal.SetFlagString(&opts.Log, "log", config)

	internal.DrawWatermark(internal.WatermarkText())

	if *versionFlag {
		printVersion()
		return
	}

	if err := internal.Setup(*opts); err != nil {
		logger.SharedLogger.Fatalf("Error during setup: %v", err)
	}

	if err := internal.Process(*opts); err != nil {
		logger.SharedLogger.Fatalf("Error during processing: %v", err)
	}

	if err := internal.CleanTemp(*opts); err != nil {
		logger.SharedLogger.Fatalf("Error cleaning temp directory: %v", err)
	}

	if err := internal.CleanEmpty(*opts); err != nil {
		logger.SharedLogger.Fatalf("Error cleaning empty directories: %v", err)
	}
}
