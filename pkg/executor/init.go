package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (e *Executor) InitCommand() error {
	statement := strings.Split("python -m venv venv", " ")
	cmd := exec.Command(statement[0], statement[1:]...)
	cmd.Dir = e.Path
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error installing on initializing venv: %v", err)
	}
	return err
}
