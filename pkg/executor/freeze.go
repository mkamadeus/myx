package executor

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func (e *Executor) FreezeCommand() error {
	statement := strings.Split(fmt.Sprintf("%s freeze > requirements.txt", path.Join(e.Path, "venv/bin/pip")), " ")
	cmd := exec.Command(statement[0], statement[1:]...)
	cmd.Dir = e.Path
	_, err := cmd.Output()
	return err
}
