package api

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/mkamadeus/myx/pkg/template/input"
	"github.com/mkamadeus/myx/pkg/template/model"
	"github.com/mkamadeus/myx/pkg/template/output"
	"github.com/mkamadeus/myx/pkg/template/pipeline"
)

//go:embed api.template
var APICode string

type APIValues struct {
	PipelineCode *pipeline.PipelineCode
	InputCode    *input.InputCode
	OutputCode   *output.OutputCode
	ModelCode    *model.ModelCode
}

func RenderAPICode(values *APIValues) (string, error) {
	t, err := template.New("api_code").Parse(APICode)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
