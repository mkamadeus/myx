package tabular

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/generator/common/model"
)

func (g *TabularGenerator) RenderModelSpec() (*generator.ModelCode, error) {
	var module model.ModelModule
	imports := make([]string, 0)

	if g.Spec.Model.Format == "keras" {
		module = &model.KerasModule{
			Path: g.Spec.Model.Path,
		}
		imports = append(imports, "from keras.models import load_model")
	} else if g.Spec.Model.Format == "onnx" {
		module = &model.ONNXModule{
			Path: g.Spec.Model.Path,
		}
		imports = append(imports, "import onnxruntime as rt")
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
		Imports:    imports,
		Session:    sessionCode,
		Prediction: predictionCode,
	}, nil

}
