package flag

import (
	"log"
	"os"
)

func OpenLogFile(opt Options) *os.File {
	file, err := os.OpenFile(LogPath(opt), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
