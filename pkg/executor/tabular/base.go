package tabular

type TabularExecutor struct {
	Imports []string
	Path    string
}

func (e *TabularExecutor) Execute() error {
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
