package spec

import "gopkg.in/yaml.v2"

type MyxSpec struct {
	Input     InputSpec     `yaml:"input"`
	Output    OutputSpec    `yaml:"output"`
	Pipeline  PipelineSpec  `yaml:"pipeline"`
	Model     ModelSpec     `yaml:"model"`
	Interface InterfaceSpec `yaml:"interface"`
}

func Parse(specBytes []byte) (*MyxSpec, error) {
	spec := &MyxSpec{}
	err := yaml.Unmarshal(specBytes, spec)
	if err != nil {
		return nil, err
	}
	return spec, nil
}
