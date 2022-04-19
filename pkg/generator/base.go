package generator

type Generator interface {
	RenderCode() (string, error)
	RenderAPISpec() (*APICode, error)
	RenderInputSpec() (*InputCode, error)
	RenderOutputSpec() (*OutputCode, error)
	RenderModelSpec() (*ModelCode, error)
	RenderPipelineSpec() (*PipelineCode, error)
}
