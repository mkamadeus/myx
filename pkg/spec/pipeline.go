package spec

type PipelineSpec []struct {
	Module   string                 `yaml:"module"`
	Metadata map[string]interface{} `yaml:"meta"`
}
