package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Verbose bool
var Output string

func New() *AppConfig {
	config := &AppConfig{
		Verbose: Verbose,
		Output:  Output,
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
