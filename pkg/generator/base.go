package generator

import (
	_ "embed"

	"github.com/mkamadeus/myx/pkg/generator/input"
	"github.com/mkamadeus/myx/pkg/generator/pipeline"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/api"
)

func RenderSpec(s *spec.MyxSpec) (string, error) {
	// input
	logger.Logger.Instance.Info("rendered input specification")
	inputCode, err := input.RenderInputSpec(s)
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Debug(inputCode)

	// output
	logger.Logger.Instance.Info("rendered output specification")
	outputCode, err := RenderOutputSpec(s)
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Debug(outputCode)

	// model
	logger.Logger.Instance.Info("rendered model specification")
	modelCode, err := RenderModelSpec(s)
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Debug(modelCode)

	// pipeline
	pipelineCode, err := pipeline.RenderPipelineSpec(s)
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Info("rendered pipeline specification")
	logger.Logger.Instance.Debug(pipelineCode)

	values := &api.APIValues{
		InputCode:    inputCode,
		OutputCode:   outputCode,
		PipelineCode: pipelineCode,
		ModelCode:    modelCode,
	}
	apiCode, err := api.RenderAPICode(values)
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Debug(apiCode)

	return apiCode, nil
}
