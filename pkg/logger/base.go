package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var Logger *logger

type logger struct {
	Level    logrus.Level
	Instance logrus.Logger
	Target   io.Writer
}

func New(level logrus.Level, target io.Writer) (*logger, error) {
	instance := logrus.New()
	instance.SetLevel(level)
	instance.SetOutput(target)

	return &logger{
		Level:    level,
		Target:   target,
		Instance: *instance,
	}, nil
}
