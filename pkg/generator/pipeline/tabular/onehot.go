package tabular

import (
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

type OneHotModule struct {
	Targets []int
	Name string
	Values []string
	NumpyTypes []string
}

func (module *OneHotModule) Run() ([]string, error) {
	result := make([]string, 0)
	for it := range module.Targets {
		values := &tabular.TabularOneHotValues{
			Index: module.Targets[it],
			Name: module.Name,
			Value: module.Values[it],
		}
		code, err := tabular.GenerateTabularOneHotCode(values)
		if err != nil {
			return nil, err
		}
		result = append(result, code...)
	}

	return result, nil
}
