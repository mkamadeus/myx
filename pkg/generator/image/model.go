package image

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/generator/common/model"
)

func (g *ImageGenerator) RenderModelSpec() (*generator.ModelCode, error) {
	var module model.ModelModule

	if g.Spec.Model.Format == "keras" {
		module = &model.KerasModule{
			Path: g.Spec.Model.Path,
		}
	} else if g.Spec.Model.Format == "onnx" {
		module = &model.ONNXModule{
			Path: g.Spec.Model.Path,
		}
	} else {
		return nil, fmt.Errorf("invalid model format %s found", g.Spec.Model.Format)
	}

	sessionCode, err := module.GetSessionCode()
	if err != nil {
		return nil, err
	}
	predictionCode, err := module.GetPredictionCode()
	if err != nil {
		return nil, err
	}
	return &generator.ModelCode{
		Session:    sessionCode,
		Prediction: predictionCode,
	}, nil

}
