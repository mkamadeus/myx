package tabular

import (
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

func DirectModule(input map[string]interface{}) *tabular.TabularNormalValues {
	return &tabular.TabularNormalValues{
		Index:     input["target"].(int),
		Name:      input["name"].(string),
		NumpyType: models.NumpyTypeMapper[input["type"].(string)],
	}
}
