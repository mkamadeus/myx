package tabular

import (
	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/logger"
)

func (g *TabularGenerator) RenderAPISpec() (*generator.APICode, error) {
	// input
	inputCode, err := g.RenderInputSpec()
	if err != nil {
		return nil, err
	}
	logger.Logger.Instance.Info("rendered input specification")
	logger.Logger.Instance.Debug(inputCode)

	// output
	outputCode, err := g.RenderOutputSpec()
	if err != nil {
		return nil, err
	}
	logger.Logger.Instance.Info("rendered output specification")
	logger.Logger.Instance.Debug(outputCode)

	// model
	modelCode, err := g.RenderModelSpec()
	if err != nil {
		return nil, err
	}
	logger.Logger.Instance.Info("rendered model specification")
	logger.Logger.Instance.Debug(modelCode)

	// pipeline
	pipelineCode, err := g.RenderPipelineSpec()
	if err != nil {
		return nil, err
	}
	logger.Logger.Instance.Info("rendered pipeline specification")
	logger.Logger.Instance.Debug(pipelineCode)

	// api code
	apiCode := &generator.APICode{
		InputCode:    inputCode,
		OutputCode:   outputCode,
		PipelineCode: pipelineCode,
		ModelCode:    modelCode,
	}
	logger.Logger.Instance.Debug(apiCode)

	return apiCode, nil
}
