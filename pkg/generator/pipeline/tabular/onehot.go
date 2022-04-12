package tabular

import (
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

func OneHotModule(input map[string]interface{}, pipelineData *spec.Pipeline) []*tabular.TabularOneHotValues {
	values := pipelineData.Metadata["values"].([]interface{})
	onehotValues := make([]string, len(values))
	for ival, val := range values {
		onehotValues[ival] = val.(string)
	}

	targets := pipelineData.Metadata["target"].([]interface{})
	onehotTargets := make([]int, len(targets))
	for itar, tar := range targets {
		onehotTargets[itar] = tar.(int)
	}

	result := make([]*tabular.TabularOneHotValues, 0)
	for ival := range onehotValues {
		result = append(result, &tabular.TabularOneHotValues{
			Index:     onehotTargets[ival],
			Name:      input["name"].(string),
			NumpyType: models.NumpyTypeMapper[input["type"].(string)],
			Value:     onehotValues[ival],
		})
	}

	return result
}
