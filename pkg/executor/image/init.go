package image

import (
	"os/exec"
	"strings"
)

func (e *ImageExecutor) InitCommand() error {
	statement := strings.Split("python -m venv venv", " ")
	cmd := exec.Command(statement[0], statement[1:]...)
	cmd.Dir = e.Path
	_, err := cmd.Output()
	return err
}
