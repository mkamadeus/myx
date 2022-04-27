package pipeline

type PipelineModule interface {
	Run() ([]string, error)
}
