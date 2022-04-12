package tabular

import (
	"fmt"
	"sort"

	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/pipeline"
)

func RenderTabularPipelineSpec(s *spec.MyxSpec) (*pipeline.PipelineCode, error) {
	logger.Logger.Instance.Debug("running in tabular input mode")

	// find sessions
	for _, p := range s.Pipeline {
		if p.Module == "preprocessing/scale" {
			values := ScaleSession(&p)
			code, err := pipeline.RenderTabularPipelineSession(values)
		}
	}

	// map input in temporary buffer
	inputMapper := make(map[int]interface{})

	logger.Logger.Instance.Info("mapping input in temporary buffer")
	for _, input := range s.Input.Metadata {
		// if input is not preprocessed
		if input["preprocessed"] == nil || input["preprocessed"] == false {
			logger.Logger.Instance.Debugf("direct input %v", input)
			inputMapper[input["target"].(int)] = DirectModule(input)
		} else {
			// else when input is preprocessed
			logger.Logger.Instance.Debugf("preprocessed input %v, detecting module", input)
			for _, p := range s.Pipeline {
				// find the preprocessing module
				if p.Metadata["for"] == input["name"] {
					logger.Logger.Instance.Debugf("using %s for input", p.Module)

					if p.Module == "preprocessing/scale" {
						inputMapper[p.Metadata["target"].(int)] = ScaleModule(input, &p)
					} else if p.Module == "preprocessing/onehot" {
						result := OneHotModule(input, &p)
						for _, r := range result {
							inputMapper[r.Index] = r
						}
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

	// get keys
	keys := make([]int, 0, len(inputMapper))
	for k := range inputMapper {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// get values
	values := make([]interface{}, 0)
	for _, k := range keys {
		values = append(values, inputMapper[k])
	}

	// variable names
	variables := make([]string, 0)
	for i := range keys {
		variables = append(variables, fmt.Sprintf("%d", i))
	}

	code, err := pipeline.RenderTabularPipelineCode(values, &pipeline.PipelineAggregationValues{
		PipelineVariables: variables,
	})
	if err != nil {
		return nil, err
	}

	return code, nil
}
