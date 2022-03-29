package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/pipeline"
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

func RenderPipelineSpec(s *spec.MyxSpec) (string, error) {

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
				inputMapper[input["target"].(int)] = &pipeline.TabularNormalValues{
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

							inputMapper[p.Metadata["target"].(int)] = &pipeline.TabularScaledValues{
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
								inputMapper[onehotTargets[ival]] = &pipeline.TabularOneHotValues{
									Index:     onehotTargets[ival],
									Name:      input["name"].(string),
									NumpyType: numpyTypeMapper[input["type"].(string)],
									Value:     onehotValues[ival],
								}
							}

						} else {
							logger.Logger.Instance.Debug("module not found")
							return "", fmt.Errorf("invalid module found")
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

		code, err := pipeline.RenderTabularPipelineCode(values)
		if err != nil {
			return "", nil
		}

		// result := make([]string, 0)
		// for _, key := range keys {
		// 	val := inputMapper[key]
		// 	pipelineType := reflect.TypeOf(val).String()

		// 	if pipelineType == "*generator.normalTabularPipeline" {
		// 		casted := val.(*normalTabularPipeline)
		// 		t, err := template.New("tabular_normal").Parse(textTemplate.TabularNormalTemplate)
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		buf := new(bytes.Buffer)
		// 		err = t.Execute(buf, &normalTabularCode{
		// 			Index:     key,
		// 			Name:      casted.Name,
		// 			NumpyType: numpyTypeMapper[strings.ToLower(casted.Type)],
		// 		})
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		result = append(result, buf.String())
		// 	} else if pipelineType == "*generator.onehotTabularPipeline" {
		// 		casted := val.(*onehotTabularPipeline)
		// 		t, err := template.New("tabular_onehot").Parse(textTemplate.TabularOnehotTemplate)
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		buf := new(bytes.Buffer)
		// 		err = t.Execute(buf, &onehotTabularCode{
		// 			Index:     key,
		// 			Name:      casted.Name,
		// 			Value:     casted.Values[casted.Index],
		// 			NumpyType: numpyTypeMapper[strings.ToLower(casted.Type)],
		// 		})
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		result = append(result, buf.String())
		// 	} else if pipelineType == "*generator.scaledTabularPipeline" {
		// 		casted := val.(*scaledTabularPipeline)
		// 		t, err := template.New("tabular_scaked").Parse(textTemplate.TabularScaledTemplate)
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		buf := new(bytes.Buffer)
		// 		err = t.Execute(buf, &scaledTabularCode{
		// 			Index:     key,
		// 			Name:      casted.Name,
		// 			NumpyType: numpyTypeMapper[strings.ToLower(casted.Type)],
		// 			Path:      casted.Path,
		// 		})
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		result = append(result, buf.String())
		// 	}

		// }

		return strings.Join(code, ""), nil

	} else if s.Input.Format == "image" {
		// TODO: image pipeline
		return "", nil
	}

	return "", fmt.Errorf("invalid input format")
}
