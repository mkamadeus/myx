package generator

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/input"
)

func RenderInputSpec(s *spec.MyxSpec) (*input.InputCode, error) {
	// input
	if s.Input.Format == "tabular" {
		logger.Logger.Instance.Debug("running in tabular input mode")

		values := make([]*input.TabularInputTypeValues, 0)
		for _, m := range s.Input.Metadata {
			values = append(values, &input.TabularInputTypeValues{
				Name: m["name"].(string),
				Type: m["type"].(string),
			})
		}

		code, err := input.RenderTabularInputCode(values, &input.TabularInputBodyValues{})
		if err != nil {
			return nil, err
		}
		return code, nil

	} else if s.Input.Format == "image" {
		// TODO: implement input spec image
	}

	return nil, fmt.Errorf("undefined input type %s", s.Input.Format)
}
