package pipeline

import (
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

type LabelModule struct {
	Path    string
	Names   []string
	Targets []int
}

func (module *LabelModule) Run() ([]string, error) {
	result := make([]string, 0)

	// make label code once
	values := &tabular.TabularLabellerValues{
		Names: module.Names,
		Path:  module.Path,
	}
	scalerCode, err := tabular.GenerateTabularLabellerCode(values)
	if err != nil {
		return nil, err
	}
	result = append(result, scalerCode...)

	// map scaler result to each
	for it := range module.Targets {
		values := &tabular.TabularLabelValues{
			Index:    module.Targets[it],
			Name:     module.Names[it],
			Position: it,
		}
		code, err := tabular.GenerateTabularLabelCode(values)
		if err != nil {
			return nil, err
		}
		result = append(result, code...)

	}

	return result, nil

}
