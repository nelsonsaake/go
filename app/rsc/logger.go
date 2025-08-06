package rsc

import (
	"io"
	"log"
	"os"
)

// SilentLogger disables all output from the default logger
type SilentLogger struct{}

func (SilentLogger) Setup() error {
	log.SetOutput(io.Discard)
	return nil
}

// VerboseLogger enables simple plain output (no timestamps, no prefixes)
type VerboseLogger struct{}

func (VerboseLogger) Setup() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(0) // no timestamp, no file info
	return nil
}
