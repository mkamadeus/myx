package config

import (
	"io"

	"github.com/sirupsen/logrus"
)

var Config *AppConfig

type AppConfig struct {
	Verbose bool
	Output  string
	Path    string
	Log     *LogConfig
}

type LogConfig struct {
	Level  logrus.Level
	Output io.Writer
}
