package model

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/models/spec"
	"github.com/mkamadeus/myx/pkg/template/model"
)

func RenderModelSpec(s *spec.MyxSpec) (*model.ModelCode, error) {
	if s.Model.Format == "keras" {
		code, err := RenderKerasModelCode(s)

		if err != nil {
			return nil, err
		}
		return code, nil
	} else if s.Model.Format == "onnx" {
		code, err := RenderONNXModelCode(s)

		if err != nil {
			return nil, err
		}
		return code, nil
	}

	return nil, fmt.Errorf("invalid model format %s found", s.Model.Format)
}
