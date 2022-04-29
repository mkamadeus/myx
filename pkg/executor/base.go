package executor

import "github.com/mkamadeus/myx/pkg/logger"

type executor interface {
	Execute() error
}

type Executor struct {
	Imports []string
	Path    string
}

func (e *Executor) Execute() error {
	logger.Logger.Instance.Debug("initializing venv")
	err := e.InitCommand()
	if err != nil {
		return err
	}
	logger.Logger.Instance.Debug("setting imports")
	err = e.ImportsCommand()
	if err != nil {
		return err
	}
	logger.Logger.Instance.Debug("freezing requirements")
	err = e.FreezeCommand()
	if err != nil {
		return err
	}
	return nil
}
