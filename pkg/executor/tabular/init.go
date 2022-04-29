package tabular

import (
	"os/exec"
	"strings"
)

func (e *TabularExecutor) InitCommand() error {
	statement := strings.Split("python -m venv venv", " ")
	cmd := exec.Command(statement[0], statement[1:]...)
	cmd.Dir = e.Path
	err := cmd.Run()
	return err
}
