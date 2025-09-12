package logs

import (
	"io"
	"log"
)

func Silent() {
	log.SetOutput(io.Discard)
}
