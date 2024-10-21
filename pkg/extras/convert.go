package extras

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ricochhet/modmanager/pkg/logger"
)

var (
	errInvalidMode       = errors.New("invalid mode: specify 'hex-to-other' or 'other-to-hex'")
	errInvalidConversion = errors.New("invalid conversion: specify 'string' or 'decimal'")
)

func NewConvert(args []string) error {
	switch args[0] {
	case "hex-to-other":
		if err := hexToOther(args[1], args[2]); err != nil {
			return err
		}
	case "other-to-hex":
		if err := otherToHex(args[1], args[2]); err != nil {
			return err
		}
	default:
		return errInvalidMode
	}

	return nil
}

func hexToOther(input string, conversionType string) error {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		return err
	}

	switch conversionType {
	case "string":
		logger.SharedLogger.Infof("String: %s", string(bytes))
	case "decimal":
		for _, b := range bytes {
			logger.SharedLogger.Infof("Decimal: %d ", b)
		}

		logger.SharedLogger.Info("\n")
	default:
		return errInvalidConversion
	}

	return nil
}

func otherToHex(input string, conversionType string) error {
	var bytes []byte

	switch conversionType {
	case "string":
		bytes = []byte(input)
	case "decimal":
		decimalStrings := strings.Split(input, ",")
		for _, ds := range decimalStrings {
			val, err := strconv.Atoi(ds)

			if err != nil || val < 0 || val > 255 {
				return fmt.Errorf("invalid decimal value: %s", ds) //nolint:err113 // wontfix
			}

			bytes = append(bytes, byte(val))
		}
	default:
		return errInvalidConversion
	}

	logger.SharedLogger.Infof("Bytes: %s\n", hex.EncodeToString(bytes))

	return nil
}
