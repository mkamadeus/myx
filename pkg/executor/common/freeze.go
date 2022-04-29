package common

import (
	"os/exec"
	"strings"
)

type CommonExecutor struct {
	Path string
}

func (e *CommonExecutor) FreezeCommand() error {
	statement := strings.Split("pip freeze > requirements.txt", " ")
	cmd := exec.Command(statement[0], statement[1:]...)
	cmd.Dir = e.Path
	_, err := cmd.Output()
	return err
}
