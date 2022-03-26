package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Verbose bool

func New() *AppConfig {
	config := &AppConfig{
		Verbose: Verbose,
	}

	// default logger config
	config.Log = &LogConfig{
		Level:  logrus.InfoLevel,
		Output: os.Stdout,
	}
	if config.Verbose {
		config.Log.Level = logrus.DebugLevel
	}

	return config
}
