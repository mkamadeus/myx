package config

import (
	"io"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var Config *AppConfig

type AppConfig struct {
	dig.In

	Verbose bool
	Path    string
	Log     *LogConfig
}

type LogConfig struct {
	dig.In

	Level  logrus.Level
	Output io.Writer
}
