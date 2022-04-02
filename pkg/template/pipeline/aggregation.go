package pipeline

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed aggregation.template
var PipelineAggregationTemplate string

type PipelineAggregationValues struct {
	PipelineVariables []string
}

func GeneratePipelineAggregationCode(values *PipelineAggregationValues) (string, error) {
	t, err := template.New("aggregation").Parse(PipelineAggregationTemplate)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		panic(err)
	}

	return buf.String(), nil
}
