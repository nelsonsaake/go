package rsc

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

// SilentLogger disables all logging
type SilentLogger struct{}

func (SilentLogger) Setup() {
	logrus.SetOutput(io.Discard)
}

// VerboseLogger enables plain output (no timestamps or levels)
type VerboseLogger struct{}

func (VerboseLogger) Setup() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		DisableSorting:         true,
		DisableQuote:           true,
		DisableColors:          true,
		PadLevelText:           false,
	})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(false)
	logrus.SetOutput(&plainWriter{})
}

// plainWriter strips all formatting and just prints the log message
type plainWriter struct{}

func (w *plainWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}
