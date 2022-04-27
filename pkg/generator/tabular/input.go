package tabular

import (
	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/template/input/tabular"
)

func (g *TabularGenerator) RenderInputSpec() (*generator.InputCode, error) {

	logger.Logger.Instance.Debug("running in tabular input mode")

	values := make([]*tabular.TabularInputTypeValues, 0)
	logger.Logger.Instance.Debug(g.Spec)
	columns := g.Spec.Input.Metadata["columns"].([]interface{})
	for _, v := range columns {
		casted := v.(map[interface{}]interface{})
		values = append(values, &tabular.TabularInputTypeValues{
			Name: casted["name"].(string),
			Type: models.BodyTypeMapper[casted["type"].(string)],
		})
	}

	typeCode, err := tabular.GenerateTabularInputTypeCode(values)
	if err != nil {
		return nil, err
	}

	bodyCode, err := tabular.GenerateTabularInputBodyCode(&tabular.TabularInputBodyValues{})
	if err != nil {
		return nil, err
	}

	return &generator.InputCode{
		Type: typeCode,
		Body: bodyCode,
	}, nil

}
