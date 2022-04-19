package tabular

import (
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

type ScaleModule struct {
	Path string
	Names []string
	Targets []int
}

func (module *ScaleModule) Run() ([]string, error) {
	result := make([]string, 0);

	// make scaler code once
	values := &tabular.TabularScalerValues{
		Names: module.Names,
		Path:  module.Path,
	}
	scalerCode, err := tabular.GenerateTabularScalerCode(values)
	if err != nil {
		return nil, err
	}
	result = append(result, scalerCode...)

	// map scaler result to each 
	for it := range module.Targets {
		values := &tabular.TabularScaledValues{
			Index:     module.Targets[it],
			Name:      module.Names[it],
			Position:  it,
		}
		code, err := tabular.GenerateTabularScaledCode(values)
		if err != nil {
			return nil, err
		}
		result = append(result, code...)

	}

	return result, nil

}