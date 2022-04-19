package image

import (
	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/logger"
)

func (g *ImageGenerator) RenderAPISpec() (*generator.APICode, error) {
	// input
	logger.Logger.Instance.Info("rendered input specification")
	inputCode, err := g.RenderInputSpec()
	if err != nil {
		return nil, err
	}
	logger.Logger.Instance.Debug(inputCode)

	// output
	logger.Logger.Instance.Info("rendered output specification")
	outputCode, err := g.RenderOutputSpec()
	if err != nil {
		return nil, err
	}
	logger.Logger.Instance.Debug(outputCode)

	// model
	logger.Logger.Instance.Info("rendered model specification")
	modelCode, err := g.RenderModelSpec()
	if err != nil {
		return nil, err
	}
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

	return apiCode, nil
}
