package image

import (
	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/template/output"
)

func (g *ImageGenerator) RenderOutputSpec() (*generator.OutputCode, error) {
	typeValues := make([]*output.OutputTypeValues, 0)
	for _, m := range g.Spec.Output {
		typeValues = append(typeValues, &output.OutputTypeValues{
			Name: m.Name,
			Type: m.Type,
		})
	}

	typeCode, err := output.GenerateOutputType(typeValues)
	if err != nil {
		return nil, err
	}
	predictionCode, err := output.GenerateOutputPrediction(&output.OutputPredictionValues{})
	if err != nil {
		return nil, err
	}

	return &generator.OutputCode{
		Type:       typeCode,
		Prediction: predictionCode,
	}, nil
}
