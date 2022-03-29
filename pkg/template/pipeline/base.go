package pipeline

import (
	"fmt"
	"reflect"
)

type PipelineCode []string

// TODO: interface{} fix
func RenderTabularPipelineCode(values []interface{}) (PipelineCode, error) {
	result := make(PipelineCode, 0)
	for _, val := range values {
		pipelineType := reflect.TypeOf(val).String()
		if pipelineType == "*pipeline.TabularNormalValues" {
			casted := val.(*TabularNormalValues)
			code, err := GenerateTabularNormalCode(casted)
			if err != nil {
				return nil, err
			}
			result = append(result, code)
		} else if pipelineType == "*pipeline.TabularOneHotValues" {
			casted := val.(*TabularOneHotValues)
			code, err := GenerateTabularOneHotCode(casted)
			if err != nil {
				return nil, err
			}
			result = append(result, code)
		} else if pipelineType == "*pipeline.TabularScaledValues" {
			casted := val.(*TabularScaledValues)
			code, err := GenerateTabularScaledCode(casted)
			if err != nil {
				return nil, err
			}
			result = append(result, code)
		} else {
			return nil, fmt.Errorf("invalid pipeline type %s found", pipelineType)
		}
	}
	return result, nil
}
