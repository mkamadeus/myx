package models

type MyxSpec struct {
	Input     InputSpec     `yaml:"input"`
	Output    OutputSpec    `yaml:"output"`
	Pipeline  PipelineSpec  `yaml:"pipeline"`
	Model     ModelSpec     `yaml:"model"`
	Interface InterfaceSpec `yaml:"interface"`
}

type InputSpec struct {
	Format   string                      `yaml:"format"`
	Metadata map[interface{}]interface{} `yaml:"meta"`
}

type OutputSpec []Output

type Output struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type ModelSpec struct {
	Format string
	Path   string
}

type PipelineSpec []Pipeline

type Pipeline struct {
	Module   string                      `yaml:"module"`
	Metadata map[interface{}]interface{} `yaml:"meta"`
}

type InterfaceSpec []Interface

type Interface struct {
	Type string `yaml:"type"`
	Port string `yaml:"port"`
}
