package spec

type PipelineSpec []struct {
	Module       string                 `yaml:"module"`
	TargetInput  string                 `yaml:"for"`
	TargetColumn []int                  `yaml:"target"`
	Metadata     map[string]interface{} `yaml:"meta"`
}
