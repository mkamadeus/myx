package generator

import (
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/output"
)

func RenderOutputSpec(s *spec.MyxSpec) (*output.OutputCode, error) {
	typeValues := make([]*output.OutputTypeValues, 0)
	for _, m := range s.Output {
		typeValues = append(typeValues, &output.OutputTypeValues{
			Name: m.Name,
			Type: m.Type,
		})
	}
	outputCode, err := output.RenderOutputCode(typeValues, &output.OutputPredictionValues{})
	if err != nil {
		return nil, err
	}
	return outputCode, nil
}
