package executor

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/mkamadeus/myx/pkg/logger"
)

func (e *Executor) ImportsCommand() error {

	for _, i := range e.Imports {
		statement := strings.Split(fmt.Sprintf("%s install %s", path.Join(e.Path, "venv/bin/pip"), i), " ")
		logger.Logger.Instance.Debug(statement)
		cmd := exec.Command(statement[0], statement[1:]...)
		cmd.Dir = e.Path
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error installing on package %s", i)
		}
	}

	return nil
}
