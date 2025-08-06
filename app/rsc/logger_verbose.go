package rsc

import (
	"github.com/sirupsen/logrus"
)

// UseLogger enables basic logging at Info level
type UseLogger struct{}

func (v UseLogger) Setup() error {
	logrus.SetLevel(logrus.InfoLevel)
	return nil
}
