package generator

import (
	_ "embed"

	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/api"
)

func RenderSpec(s *spec.MyxSpec) error {
	// input
	logger.Logger.Instance.Info("rendered input specification")
	inputSpec, err := RenderInputSpec(s)
	if err != nil {
		return err
	}
	logger.Logger.Instance.Debug(inputSpec)

	// output
	logger.Logger.Instance.Info("rendered output specification")
	outputSpec, err := RenderOutputSpec(s)
	if err != nil {
		return err
	}
	logger.Logger.Instance.Debug(outputSpec)

	// model
	logger.Logger.Instance.Info("rendered model specification")
	modelSpec, err := RenderModelSpec(s)
	if err != nil {
		return err
	}
	logger.Logger.Instance.Debug(modelSpec)

	// pipeline
	pipelineSpec, err := RenderPipelineSpec(s)
	if err != nil {
		return err
	}
	logger.Logger.Instance.Info("rendered pipeline specification")
	logger.Logger.Instance.Debug(pipelineSpec)

	values := &api.APIValues{
		InputCode:    inputSpec,
		OutputCode:   outputSpec,
		PipelineCode: pipelineSpec,
		ModelCode:    modelSpec,
	}
	apiCode, err := api.RenderAPICode(values)
	if err != nil {
		return err
	}
	logger.Logger.Instance.Debug(apiCode)

	return nil
}
