package generator

import (
	"fmt"
	"reflect"

	"github.com/mkamadeus/myx/pkg/spec"
)

type normalTabularPipeline struct {
	Name string
	Type string
}

type scaledTabularPipeline struct {
	Name string
	Type string
	Path string
}

type onehotTabularPipeline struct {
	Name   string
	Type   string
	Values []string
	Index  int
}

func RenderPipelineSpec(s *spec.MyxSpec) (string, error) {

	// pipeline
	if s.Input.Format == "tabular" {
		// t, err = template.New("pipeline").Parse(textTemplate.InputTemplate)
		inputMapper := make(map[int]interface{})
		fmt.Println("pisang")
		for _, input := range s.Input.Metadata {
			fmt.Println(input)
			// if input is not preprocessed
			if input["preprocessed"] == nil || input["preprocessed"] == false {
				fmt.Println("lol")
				inputMapper[input["target"].(int)] = &normalTabularPipeline{
					Name: input["name"].(string),
					Type: input["type"].(string),
				}
			} else {
				// else when input is preprocessed
				fmt.Println("lol2")
				for _, pipeline := range s.Pipeline {
					// find the preprocessing module
					if pipeline.Metadata["for"] == input["name"] {
						if pipeline.Module == "preprocessing/scale" {

							inputMapper[pipeline.Metadata["target"].(int)] = &scaledTabularPipeline{
								Name: input["name"].(string),
								Type: input["type"].(string),
								Path: fmt.Sprintf("%v", pipeline.Metadata["path"]),
							}
						} else if pipeline.Module == "preprocessing/onehot" {
							values := pipeline.Metadata["values"].([]interface{})
							onehotValues := make([]string, len(values))
							for ival, val := range values {
								onehotValues[ival] = val.(string)
							}
							fmt.Println(onehotValues)

							targets := pipeline.Metadata["target"].([]interface{})
							onehotTargets := make([]int, len(targets))
							for itar, tar := range targets {
								onehotTargets[itar] = tar.(int)
							}
							fmt.Println(onehotTargets)

							for ival := range onehotValues {
								inputMapper[onehotTargets[ival]] = &onehotTabularPipeline{
									Name:   input["name"].(string),
									Type:   input["type"].(string),
									Values: onehotValues,
									Index:  ival,
								}
							}

						} else {
							return "", fmt.Errorf("invalid module found")
						}

					}
				}

			}

		}
		fmt.Println("pisang")
		fmt.Println(inputMapper)
		for pos, val := range inputMapper {
			fmt.Printf("%d -> %s\n", pos, reflect.TypeOf(val))
		}

		return "pipeline specccc", nil

	} else if s.Input.Format == "image" {
		// TODO: image pipeline
		return "", nil
	}

	return "", fmt.Errorf("invalid input format")
}
