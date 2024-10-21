package extras

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/simplefs"
	"github.com/ricochhet/simpleutil"
)

var (
	errSpecifyPathAndBytes  = errors.New("specify both a file path and the bytes to search for")
	errBytesMustMatchLength = errors.New("replacement bytes length must match find bytes length")
	errPositionToInteger    = errors.New("error converting position to integer")
)

func NewPatch(args []string) error {
	if args[0] == "" || args[1] == "" {
		return errSpecifyPathAndBytes
	}

	if simplefs.Exists(args[0]) && simplefs.Exists(args[1]) {
		data, err := ReadPatchTable(args[1])
		if err != nil {
			return err
		}

		for _, bytes := range data.Bytes {
			//nolint:lll // wontfix
			if err := findAndReplaceBytes(args[0], strings.ReplaceAll(bytes.Find, " ", ""), strings.ReplaceAll(bytes.Replace, " ", ""), bytes.Position); err != nil {
				return err
			}
		}

		return nil
	}

	if err := simpleutil.CheckArgumentCount(args, 4); err != nil {
		return err
	}

	if err := findAndReplaceBytes(args[0], args[1], args[2], args[3]); err != nil {
		return err
	}

	return nil
}

//nolint:cyclop // wontfix
func findAndReplaceBytes(fileName, searchBytes, replacementBytes, position string) error {
	findBytes, err := simpleutil.HexStringToBytes(searchBytes)
	if err != nil {
		return err
	}

	var replaceWith []byte
	if replacementBytes != "" {
		replaceWith, err = simpleutil.HexStringToBytes(replacementBytes)
		if err != nil {
			return err
		}
	}

	content, err := simplefs.ReadFile(fileName)
	if err != nil {
		return err
	}

	indices := simpleutil.FindAllByteOccurrences(content, findBytes)
	if len(indices) == 0 {
		logger.SharedLogger.Info("No occurrences found.")
		return nil
	}

	logger.SharedLogger.Infof("Found occurrences at positions: %v", indices)

	if len(replaceWith) > 0 { //nolint:nestif // wontfix
		if len(findBytes) != len(replaceWith) {
			return errBytesMustMatchLength
		}

		newPosition := 0

		if position != "" {
			s, err := strconv.Atoi(position)
			if err != nil {
				return errPositionToInteger
			}

			newPosition = s
		}

		modifiedContent := simpleutil.ReplaceByteOccurrences(content, findBytes, replaceWith, newPosition)

		if err := simplefs.WriteFile(fileName, modifiedContent, os.ModePerm); err != nil {
			return err
		}

		logger.SharedLogger.Info("Bytes replaced successfully.")
	}

	return nil
}
