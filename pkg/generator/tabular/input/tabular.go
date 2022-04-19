package input

import (
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/models/spec"
	"github.com/mkamadeus/myx/pkg/template/input"
	"github.com/mkamadeus/myx/pkg/template/input/tabular"
)

func RenderTabularInputSpec(s *spec.MyxSpec) (*input.InputCode, error) {
	logger.Logger.Instance.Debug("running in tabular input mode")

	values := make([]*tabular.TabularInputTypeValues, 0)
	for _, m := range s.Input.Metadata {
		values = append(values, &tabular.TabularInputTypeValues{
			Name: m["name"].(string),
			Type: models.BodyTypeMapper[m["type"].(string)],
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

	return &input.InputCode{
		Type: typeCode,
		Body: bodyCode,
	}, nil
}
