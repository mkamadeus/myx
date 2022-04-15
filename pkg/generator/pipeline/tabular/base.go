package tabular

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/pipeline"
)

type TabularPipelineModule interface {
	Run() ([]string, error)
}

func RenderTabularPipelineSpec(s *spec.MyxSpec) (*pipeline.PipelineCode, error) {
	logger.Logger.Instance.Debug("running in tabular input mode")

	// map input in temporary buffer
	inputMapper := make([]TabularPipelineModule, 0)

	logger.Logger.Instance.Info("mapping input in temporary buffer")
	for _, input := range s.Input.Metadata {
		// if input is not preprocessed
		if input["preprocessed"] == nil || input["preprocessed"] == false {
			logger.Logger.Instance.Debugf("direct input %v", input)

			// make module for direct input, run
			module := &DirectModule{
				Target: input["target"].(int),
			}
			inputMapper = append(inputMapper, module)
		} else {
			// else when input is preprocessed
			logger.Logger.Instance.Debugf("preprocessed input %v, detecting module", input)
			for _, p := range s.Pipeline {

				// find the preprocessing module
				if p.Metadata["for"] == input["name"] {
					logger.Logger.Instance.Debugf("using %s for input", p.Module)

					if p.Module == "preprocessing/scale" {

						// add module to mapper
						module := &ScaleModule{
							Names: p.Metadata["for"].([]string),
							Targets: p.Metadata["target"].([]int),
							Path: p.Metadata["path"].(string),
						}
						inputMapper = append(inputMapper, module)
					} else if p.Module == "preprocessing/onehot" {
						// add module to mapper
						module := &OneHotModule{
							Name: p.Metadata["for"].(string),
							Targets: p.Metadata["target"].([]int),
							Values: p.Metadata["values"].([]string),
						}
						inputMapper = append(inputMapper, module)
					} else {
						logger.Logger.Instance.Debug("module not found")
						return nil, fmt.Errorf("invalid module found")
					}

				}
			}

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
	}

	

	return code, nil
}
