package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ricochhet/minicommon/util"
	"github.com/ricochhet/modmanager/info"
	"github.com/ricochhet/modmanager/pkg/extras"
	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/modmanager/pkg/reepak"
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

//nolint:gocognit,gocyclo,cyclop // wontfix
func main() {
	logger.SharedLogger = logger.NewLogger(logger.InfoLevel, os.Stdout, log.Lshortfile|log.LstdFlags)

	versionFlag := flag.Bool("v", false, "Print the current version")

	flag.Parse()

	info.DrawWatermark(info.WatermarkText())

	if *versionFlag {
		printVersion()
		return
	}

	if _, err := util.NewCommand(os.Args, "help", 0); err == nil {
		logger.SharedLogger.Info("conv <hex-to-other/other-to-hex[conv mode]> <input> <string/decimal[conv type]>")
		logger.SharedLogger.Info("patch <file> <bytesToFind> <bytesToReplace> <replaceAtOccurrence>")
		logger.SharedLogger.Info("hash <file> <murmur3x64_128hash/murmur3x86_128hash/murmur3x86_32hash/crc64/crc32/sha512/sha256/sha1/md5[mode]>")
		logger.SharedLogger.Info("diff <folderA> <folderB>")
		logger.SharedLogger.Info("reepak:")
		logger.SharedLogger.Info("\tpak <folder> <output> <0/1[embed data]>")
		logger.SharedLogger.Info("\tunpak <folder> <output> <0/1[embed data]>")
		logger.SharedLogger.Info("\tcompress <file>")
		logger.SharedLogger.Info("\tdecompress <file>")
		os.Exit(1)
	}

	if args, err := util.NewCommand(os.Args, "conv", 3); err == nil {
		if err := extras.NewConvert(args); err != nil {
			logger.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}

	if args, err := util.NewCommand(os.Args, "patch", 2); err == nil {
		if err := extras.NewPatch(args); err != nil {
			logger.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}

	if args, err := util.NewCommand(os.Args, "hash", 1); err == nil {
		if err := extras.NewHash(args); err != nil {
			logger.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}

	if args, err := util.NewCommand(os.Args, "diff", 2); err == nil {
		if err := extras.NewDiff(args); err != nil {
			logger.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}

	if args, err := util.NewCommand(os.Args, "pak", 3); err == nil {
		selection, err := strconv.Atoi(args[2])
		if err != nil {
			logger.SharedLogger.Fatalf("Error converting string to integer: %s", err)
		}

		if err := reepak.ProcessDirectory(args[0], args[1], selection != 0); err != nil {
			logger.SharedLogger.Fatalf("Error processing directory: %s", err)
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}

	if args, err := util.NewCommand(os.Args, "unpak", 3); err == nil {
		selection, err := strconv.Atoi(args[2])
		if err != nil {
			logger.SharedLogger.Fatalf("Error converting string to integer: %s", err)
		}

		if err := reepak.ExtractDirectory(args[0], args[1], selection != 0); err != nil {
			logger.SharedLogger.Fatalf("Error extracting directory: %s", err)
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}

	if args, err := util.NewCommand(os.Args, "compress", 1); err == nil {
		if err := reepak.CompressPakData(args[0]); err != nil {
			logger.SharedLogger.Fatalf("Error compressing pak data: %s", err)
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}

	if args, err := util.NewCommand(os.Args, "decompress", 1); err == nil {
		if err := reepak.DecompressPakData(args[0]); err != nil {
			logger.SharedLogger.Fatalf("Error decompressing pak data: %s", err)
		}
	} else if !errors.Is(err, util.ErrNoFunctionName) {
		logger.SharedLogger.Fatal(err.Error())
	}
}
