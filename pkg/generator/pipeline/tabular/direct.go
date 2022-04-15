package tabular

import (
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

type DirectModule struct{
	Target int
	Name string
	NumpyType string
}

func (module *DirectModule) Run() ([]string, error) {
	values := &tabular.TabularNormalValues{
		Index: module.Target,
		Name: module.Name,
		NumpyType: module.NumpyType,
	}
	code, err := tabular.GenerateTabularNormalCode(values)
	if err != nil {
		return nil, err
	}
	return code, err
}
