package extras

import (
	"crypto/md5"  //nolint:gosec // wontfix
	"crypto/sha1" //nolint:gosec // wontfix
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"math"
	"strconv"

	"github.com/ricochhet/modmanager/pkg/logger"
	"github.com/ricochhet/simplecrypto"
)

var errSeedToInt = errors.New("error converting seed to integer")

//nolint:gocyclo,cyclop // wontfix
func NewHash(args []string) error {
	seed := math.MaxUint32
	hashType := "md5"

	if len(args) > 1 {
		hashType = args[1]
	}

	if len(args) > 2 {
		if args[2] != "" {
			s, err := strconv.Atoi(args[2])
			if err != nil {
				return errSeedToInt
			}

			seed = s
		}
	}

	switch hashType {
	case "murmur3x64_128hash":
		logger.SharedLogger.Infof("%d", simplecrypto.Murmur3X64_128Hash(seed, args[0]))
	case "murmur3x86_128hash":
		logger.SharedLogger.Infof("%d", simplecrypto.Murmur3X86_128Hash(seed, args[0]))
	case "murmur3x86_32hash":
		logger.SharedLogger.Infof("%d", simplecrypto.Murmur3X86_32Hash(seed, args[0]))
	case "crc64":
		if err := newHash(args[0], crc64.New(crc64.MakeTable(crc32.IEEE))); err != nil {
			return err
		}
	case "crc32":
		if err := newHash(args[0], crc32.New(crc32.IEEETable)); err != nil {
			return err
		}
	case "sha512":
		if err := newHash(args[0], sha512.New()); err != nil {
			return err
		}
	case "sha256":
		if err := newHash(args[0], sha256.New()); err != nil {
			return err
		}
	case "sha1":
		if err := newHash(args[0], sha1.New()); err != nil { //nolint:gosec // wontfix
			return err
		}
	case "md5":
		if err := newHash(args[0], md5.New()); err != nil { //nolint:gosec // wontfix
			return err
		}
	default:
		if err := newHash(args[0], md5.New()); err != nil { //nolint:gosec // wontfix
			return err
		}
	}

	return nil
}

func newHash(filePath string, hash hash.Hash) error {
	s, err := simplecrypto.NewHash(filePath, hash)
	if err != nil {
		return err
	}

	logger.SharedLogger.Infof("%s", s)

	return nil
}
