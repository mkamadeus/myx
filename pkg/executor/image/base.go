package image

type ImageExecutor struct {
	Imports []string
	Path    string
}

func (e *ImageExecutor) Execute() error {
	err := e.InitCommand()
	if err != nil {
		return err
	}
	err = e.ImportsCommand()
	if err != nil {
		return err
	}
	return nil
}
