package rsc

import (
	"io"
	"log"
	"os"
)

// SilentLogger disables all output from the default logger
type SilentLogger struct{}

func (SilentLogger) Setup() {
	log.SetOutput(io.Discard)
}

// VerboseLogger enables simple plain output (no timestamps, no prefixes)
type VerboseLogger struct{}

func (VerboseLogger) Setup() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0) // no timestamp, no file info
}
