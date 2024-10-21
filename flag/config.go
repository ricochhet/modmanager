package flag

import (
	"log"
	"os"
	"strings"
)

func OpenConfigFile(opt Options) *os.File {
	file, err := os.OpenFile(ConfigPath(opt), os.O_CREATE|os.O_RDONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func MapConfigFile(lines []string) (map[string]string, error) {
	keyvalues := map[string]string{}

	for _, line := range lines {
		keyvalue := strings.SplitN(line, "=", 2)

		if len(keyvalue) != 2 {
			continue // skip
		}

		keyvalues[keyvalue[0]] = keyvalue[1]
	}

	return keyvalues, nil
}
