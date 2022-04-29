package image

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func (e *ImageExecutor) ImportsCommand() error {
	for _, i := range e.Imports {
		statement := strings.Split(fmt.Sprintf("pip install %s", i), " ")
		cmd := exec.Command(statement[0], statement[1:]...)
		cmd.Path = path.Join(e.Path, "bin/pip")
		cmd.Dir = e.Path
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
