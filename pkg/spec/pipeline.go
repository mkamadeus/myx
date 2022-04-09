package spec

type PipelineSpec []Pipeline

type Pipeline struct {
	Module   string                 `yaml:"module"`
	Metadata map[string]interface{} `yaml:"meta"`
}
