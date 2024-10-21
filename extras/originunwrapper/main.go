package main

import (
	"flag"
	"log"
	"os"

	"github.com/ricochhet/modmanager/internal"
	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/modmanager/pkg/originunwrapper"
	"github.com/ricochhet/readwrite"
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
	logger.SharedLogger = logger.NewLogger(logger.InfoLevel, os.Stdout, log.Lshortfile|log.LstdFlags)

	versionFlag := flag.Bool("v", false, "Print the current version")

	inputPath := flag.String("input", "", "input path")
	outputPath := flag.String("output", "", "output path")
	dlfKey := flag.String("key", "", "dlf key")
	version := flag.String("version", "", "hash version")
	getDlfKey := flag.String("dlf-path", "", "dlf path")
	addDll := flag.Bool("add-dll", false, "true / false")

	flag.Parse()

	internal.DrawWatermark(internal.WatermarkText())

	if *versionFlag {
		printVersion()
		return
	}

	if len(*inputPath) == 0 || len(*outputPath) == 0 || len(*dlfKey) == 0 || len(*version) == 0 {
		logger.SharedLogger.Fatal("Error: not enough arguments specified")
	}

	fileMap, err := readwrite.Open(*inputPath)
	if err != nil {
		logger.SharedLogger.Fatalf("Error opening file: %s", err)
	}

	originunwrapper.Unwrap(fileMap, *version, *getDlfKey, *dlfKey, *outputPath, *addDll)
}
