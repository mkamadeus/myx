package generator

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/model"
)

func RenderModelSpec(s *spec.MyxSpec) (*model.ModelCode, error) {
	if s.Model.Format == "keras" {
		kerasCode, err := model.RenderKerasModelCode(&model.ModelKerasBaseValues{
			Path: s.Model.Path,
		}, &model.ModelKerasPredictionValues{})
		if err != nil {
			return nil, err
		}

		return kerasCode, nil
	} else if s.Model.Format == "onnx" {
		onnxCode, err := model.RenderONNXModelCode(&model.ModelONNXBaseValues{
			Path: s.Model.Path,
		}, &model.ModelONNXPredictionValues{})
		if err != nil {
			return nil, err
		}

		return onnxCode, nil
	}

	return nil, fmt.Errorf("invalid model format %s found", s.Model.Format)
}
