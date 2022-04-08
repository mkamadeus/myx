package pipeline

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed aggregation.template
var PipelineAggregationTemplate string

type PipelineAggregationValues struct {
	PipelineVariables []string
}

func GeneratePipelineAggregationCode(values *PipelineAggregationValues) ([]string, error) {
	t, err := template.New("aggregation").Parse(PipelineAggregationTemplate)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		panic(err)
	}

	return utils.ClearEmptyString(strings.Split(buf.String(), "\n")), nil
}
