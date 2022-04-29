package image

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/generator/image/pipeline"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/template/pipeline/image"
)

func (g *ImageGenerator) RenderPipelineSpec() (*generator.PipelineCode, error) {
	logger.Logger.Instance.Debug("running in image pipeline mode")

	pipelines := make([]pipeline.ImagePipelineModule, 0)
	imports := make([]string, 0)

	// initial image read
	pipelines = append(pipelines, &pipeline.InitialReadModule{})
	imports = append(imports, "import io")
	imports = append(imports, "from PIL import Image")

	// for each module
	for _, p := range g.Spec.Pipeline {
		logger.Logger.Instance.Debug(p)
		if p.Module == "preprocessing/image/resize" {
			pipelines = append(pipelines, &pipeline.ResizeModule{
				Width:  p.Metadata["width"].(int),
				Height: p.Metadata["height"].(int),
			})
		} else {
			return nil, fmt.Errorf("invalid module")
		}
	}

	pipelineCodes := make([]string, 0)
	for _, p := range pipelines {
		moduleCode, err := p.Run()
		if err != nil {
			return nil, err
		}
		pipelineCodes = append(pipelineCodes, moduleCode...)
	}

	aggregationCode, err := image.GenerateImageAggregationCode(&image.ImageAggregationValues{
		Width:   g.Spec.Input.Metadata["width"].(int),
		Height:  g.Spec.Input.Metadata["height"].(int),
		Channel: g.Spec.Input.Metadata["channel"].(int),
	})
	if err != nil {
		return nil, err
	}

	return &generator.PipelineCode{
		Imports: imports,
		Pipelines:   pipelineCodes,
		Aggregation: aggregationCode,
	}, nil
}
