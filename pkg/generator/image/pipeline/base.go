package pipeline

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/template/pipeline/image"
)

type ImagePipelineModule interface {
	Run() ([]string, error)
}

// TODO: image pipeline
func RenderImagePipelineSpec(s *models.MyxSpec) (*generator.PipelineCode, error) {
	logger.Logger.Instance.Debug("running in image input mode")

	// map input in temporary buffer, save targets
	inputMapper := make([]ImagePipelineModule, 0)

	logger.Logger.Instance.Info("mapping input in temporary buffer")

	// initial read
	inputMapper = append(inputMapper, &InitialReadModule{})

	// for each module
	for _, p := range s.Pipeline {
		logger.Logger.Instance.Debug(p)

		// find the preprocessing module
		if p.Module == "preprocessing/image/resize" {
			// add module to mapper
			names := make([]string, 0)
			for _, n := range p.Metadata["for"].([]interface{}) {
				names = append(names, n.(string))
			}
			targets := make([]int, 0)
			for _, t := range p.Metadata["target"].([]interface{}) {
				targets = append(targets, t.(int))
			}

			module := &ResizeModule{
				Width:  p.Metadata["width"].(int),
				Height: p.Metadata["height"].(int),
			}
			inputMapper = append(inputMapper, module)
		} else {
			logger.Logger.Instance.Debug("module not found")
			return nil, fmt.Errorf("invalid module found")
		}
	}

	// map buffer to actual code
	logger.Logger.Instance.Info("mapping buffer to code")

	pipelineCodes := make([]string, 0)
	for _, mapper := range inputMapper {
		c, err := mapper.Run()
		if err != nil {
			return nil, err
		}

		pipelineCodes = append(pipelineCodes, c...)
		logger.Logger.Instance.Debug(pipelineCodes)
	}

	// aggregation code

	aggregationCode, err := image.GenerateImageAggregationCode(&image.ImageAggregationValues{
		Width:   s.Input.Metadata["width"].(int),
		Height:  s.Input.Metadata["height"].(int),
		Channel: s.Input.Metadata["channel"].(int),
	})
	if err != nil {
		return nil, err
	}

	return &generator.PipelineCode{
		Pipelines:   pipelineCodes,
		Aggregation: aggregationCode,
	}, nil

}
