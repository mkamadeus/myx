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
	for _, v := range g.Spec.Input.Metadata {
		casted := v.(map[string]string)
		values = append(values, &tabular.TabularInputTypeValues{
			Name: casted["name"],
			Type: models.BodyTypeMapper[casted["type"]],
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
