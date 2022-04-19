package generator

import (
	_ "embed"

	"github.com/mkamadeus/myx/pkg/generator/model"
	"github.com/mkamadeus/myx/pkg/generator/output"
	"github.com/mkamadeus/myx/pkg/generator/pipeline"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models/spec/tabular"
	"github.com/mkamadeus/myx/pkg/template/api"
)

type TabularGenerator struct {
	Spec tabular.TabularSpec
}

func (s *tabular.TabularSpec) RenderSpec() (string, error) {
	// input
	logger.Logger.Instance.Info("rendered input specification")
	inputCode, err := s.RenderInputSpec()
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Debug(inputCode)

	// output
	logger.Logger.Instance.Info("rendered output specification")
	outputCode, err := output.RenderOutputSpec(s)
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Debug(outputCode)

	// model
	logger.Logger.Instance.Info("rendered model specification")
	modelCode, err := model.RenderModelSpec(s)
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

	// api code
	values := &api.TabularAPIValues{
		InputCode:    inputCode,
		OutputCode:   outputCode,
		PipelineCode: pipelineCode,
		ModelCode:    modelCode,
	}
	apiCode, err := values.Render()
	if err != nil {
		return "", err
	}
	logger.Logger.Instance.Debug(apiCode)

	return apiCode, nil
}
