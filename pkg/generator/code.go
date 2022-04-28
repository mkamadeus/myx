package generator

type APICode struct {
	PipelineCode *PipelineCode
	InputCode    *InputCode
	OutputCode   *OutputCode
	ModelCode    *ModelCode
}

type PipelineCode struct {
	Imports     []string
	Pipelines   []string
	Aggregation []string
}

type InputCode struct {
	Imports []string
	Type    []string
	Body    []string
}

type OutputCode struct {
	Imports    []string
	Type       []string
	Prediction []string
}

type ModelCode struct {
	Imports    []string
	Session    []string
	Prediction []string
}
