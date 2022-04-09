package input

import (
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/input"
)

func RenderTabularInputSpec(s *spec.MyxSpec) (*input.InputCode, error) {
	logger.Logger.Instance.Debug("running in tabular input mode")

	values := make([]*input.TabularInputTypeValues, 0)
	for _, m := range s.Input.Metadata {
		values = append(values, &input.TabularInputTypeValues{
			Name: m["name"].(string),
			Type: models.BodyTypeMapper[m["type"].(string)],
		})
	}

	code, err := input.RenderTabularInputCode(values, &input.TabularInputBodyValues{})
	if err != nil {
		return nil, err
	}
	return code, nil
}
