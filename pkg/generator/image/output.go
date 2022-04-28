package image

import (
	"fmt"
	"strings"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/template/output"
)

func (g *ImageGenerator) RenderOutputSpec() (*generator.OutputCode, error) {

	imports := make([]string, 0)
	typeValues := make([]*output.OutputTypeValues, 0)
	for _, m := range g.Spec.Output {
		value := &output.OutputTypeValues{
			Name: m.Name,
			Type: m.Type,
		}

		// check array type, duplicates will be thrown away later
		if strings.HasSuffix(value.Type, "[]") {
			imports = append(imports, "from typing import List")
			value.Type = fmt.Sprintf("List[%s]", models.BodyTypeMapper[value.Type[0:len(value.Type)-2]])
		} else {
			value.Type = models.BodyTypeMapper[value.Type]
		}

		typeValues = append(typeValues, value)
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
		Imports:    imports,
		Type:       typeCode,
		Prediction: predictionCode,
	}, nil
}
