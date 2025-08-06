package rsc

import (
	"io"

	"github.com/sirupsen/logrus"
)

// SilentLogger disables all logging
type SilentLogger struct{}

func (s SilentLogger) Setup() error {
	logrus.SetOutput(io.Discard)
	return nil
}
