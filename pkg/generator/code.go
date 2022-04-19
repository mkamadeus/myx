package generator

type APICode struct {
	PipelineCode *PipelineCode
	InputCode    *InputCode
	OutputCode   *OutputCode
	ModelCode    *ModelCode
}

type PipelineCode struct {
	Pipelines   []string
	Aggregation []string
}

type InputCode struct {
	Type []string
	Body []string
}

type OutputCode struct {
	Type       []string
	Prediction []string
}

type ModelCode struct {
	Session    []string
	Prediction []string
}
