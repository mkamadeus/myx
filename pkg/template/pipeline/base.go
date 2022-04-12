package pipeline

import (
	"fmt"
	"reflect"

	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

type PipelineCode struct {
	Pipelines   [][]string
	Aggregation []string
}
type PipelineSessionCode struct {
	Sessions [][]string
}

// TODO: interface{} fix
func RenderTabularPipelineCode(pipelineValues []interface{}, aggregationValues *PipelineAggregationValues) (*PipelineCode, error) {
	pipelines := make([][]string, 0)
	for _, val := range pipelineValues {
		pipelineType := reflect.TypeOf(val).String()
		if pipelineType == "*tabular.TabularNormalValues" {
			casted := val.(*tabular.TabularNormalValues)
			code, err := tabular.GenerateTabularNormalCode(casted)
			if err != nil {
				return nil, err
			}
			pipelines = append(pipelines, code)
		} else if pipelineType == "*tabular.TabularOneHotValues" {
			casted := val.(*tabular.TabularOneHotValues)
			code, err := tabular.GenerateTabularOneHotCode(casted)
			if err != nil {
				return nil, err
			}
			pipelines = append(pipelines, code)
		} else if pipelineType == "*tabular.TabularScaledValues" {
			casted := val.(*tabular.TabularScaledValues)
			code, err := tabular.GenerateTabularScaledCode(casted)
			if err != nil {
				return nil, err
			}
			pipelines = append(pipelines, code)
		} else {
			return nil, fmt.Errorf("invalid pipeline type %s found", pipelineType)
		}
	}

	aggregation, err := GeneratePipelineAggregationCode(aggregationValues)
	if err != nil {
		return nil, err
	}

	return &PipelineCode{
		Pipelines:   pipelines,
		Aggregation: aggregation,
	}, nil
}

func RenderTabularPipelineSession(sessionValues []interface{}) (*PipelineSessionCode, error) {
	sessions := make([][]string, 0)
	for _, val := range sessionValues {
		sessionType := reflect.TypeOf(val).String()
		if sessionType == "*tabular.TabularScalerValues" {
			casted := val.(*tabular.TabularScalerValues)
			code, err := tabular.GenerateTabularScalerCode(casted)
			if err != nil {
				return nil, err
			}
			sessions = append(sessions, code)
		} else {
			return nil, fmt.Errorf("invalid session type %s found", sessionType)
		}
	}

	return &PipelineSessionCode{
		Sessions: sessions,
	}, nil
}
