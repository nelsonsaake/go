package app

import (
	"io"
	"log"
	"os"
)

func Silent() {
	log.SetOutput(io.Discard)
}

func Verbose() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0) // no timestamp, no file info
}
