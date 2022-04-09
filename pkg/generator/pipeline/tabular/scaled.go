package tabular

import (
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

func ScaleModule(input map[string]interface{}, pipelineData *spec.Pipeline) *tabular.TabularScaledValues {
	return &tabular.TabularScaledValues{
		Index:     pipelineData.Metadata["target"].(int),
		Name:      input["name"].(string),
		NumpyType: models.NumpyTypeMapper[input["type"].(string)],
	}
}
