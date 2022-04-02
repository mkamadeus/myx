package generator

import (
	"fmt"
	"sort"

	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/pipeline"
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

var numpyTypeMapper = map[string]string{
	"float":       "np.float32",
	"int":         "np.int_",
	"categorical": "np.int_",
}

type normalTabularPipeline struct {
	Name string
	Type string
}

type normalTabularCode struct {
	Index     int
	Name      string
	NumpyType string
}

type scaledTabularPipeline struct {
	Name string
	Type string
	Path string
}

type scaledTabularCode struct {
	Index     int
	Name      string
	Path      string
	NumpyType string
}

type onehotTabularPipeline struct {
	Name   string
	Type   string
	Values []string
	Index  int
}

type onehotTabularCode struct {
	Index     int
	Name      string
	Value     string
	NumpyType string
}

func RenderPipelineSpec(s *spec.MyxSpec) (*pipeline.PipelineCode, error) {

	// pipeline
	if s.Input.Format == "tabular" {
		logger.Logger.Instance.Debug("running in tabular input mode")

		// map input in temporary buffer
		inputMapper := make(map[int]interface{})

		logger.Logger.Instance.Info("mapping input in temporary buffer")
		for _, input := range s.Input.Metadata {
			// if input is not preprocessed
			if input["preprocessed"] == nil || input["preprocessed"] == false {
				logger.Logger.Instance.Debugf("direct input %v", input)
				inputMapper[input["target"].(int)] = &tabular.TabularNormalValues{
					Index:     input["target"].(int),
					Name:      input["name"].(string),
					NumpyType: numpyTypeMapper[input["type"].(string)],
				}
			} else {
				// else when input is preprocessed
				logger.Logger.Instance.Debugf("preprocessed input %v, detecting module", input)
				for _, p := range s.Pipeline {
					// find the preprocessing module
					if p.Metadata["for"] == input["name"] {
						logger.Logger.Instance.Debugf("using %s for input", p.Module)
						if p.Module == "preprocessing/scale" {

							inputMapper[p.Metadata["target"].(int)] = &tabular.TabularScaledValues{
								Index:     p.Metadata["target"].(int),
								Name:      input["name"].(string),
								NumpyType: numpyTypeMapper[input["type"].(string)],
								Path:      fmt.Sprintf("%v", p.Metadata["path"]),
							}
						} else if p.Module == "preprocessing/onehot" {
							values := p.Metadata["values"].([]interface{})
							onehotValues := make([]string, len(values))
							for ival, val := range values {
								onehotValues[ival] = val.(string)
							}

							targets := p.Metadata["target"].([]interface{})
							onehotTargets := make([]int, len(targets))
							for itar, tar := range targets {
								onehotTargets[itar] = tar.(int)
							}

							for ival := range onehotValues {
								inputMapper[onehotTargets[ival]] = &tabular.TabularOneHotValues{
									Index:     onehotTargets[ival],
									Name:      input["name"].(string),
									NumpyType: numpyTypeMapper[input["type"].(string)],
									Value:     onehotValues[ival],
								}
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
			variables = append(variables, fmt.Sprintf("%d", i+1))
		}

		code, err := pipeline.RenderTabularPipelineCode(values, &pipeline.PipelineAggregationValues{
			PipelineVariables: variables,
		})
		if err != nil {
			return nil, err
		}

		return code, nil

	} else if s.Input.Format == "image" {
		// TODO: image pipeline
		return nil, nil
	}

	return nil, fmt.Errorf("invalid input format")
}
