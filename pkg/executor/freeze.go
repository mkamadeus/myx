package executor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/mkamadeus/myx/pkg/logger"
)

func (e *Executor) FreezeCommand() error {
	var buf bytes.Buffer
	statement := strings.Split(fmt.Sprintf("%s freeze", path.Join(e.Path, "venv/bin/pip")), " ")
	cmd := exec.Command(statement[0], statement[1:]...)
	logger.Logger.Instance.Debug(statement)
	cmd.Dir = e.Path
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to freeze dependencies")
	}
	f, err := os.Create(path.Join(e.Path, "requirements.txt"))
	logger.Logger.Instance.Debug(buf.String())
	if err != nil {
		return err
	}

	_, err = f.WriteString(buf.String())
	return err
}
