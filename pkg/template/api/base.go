package api

import (
	"bytes"
	_ "embed"
	"html/template"

	"github.com/mkamadeus/myx/pkg/template/input"
	"github.com/mkamadeus/myx/pkg/template/model"
	"github.com/mkamadeus/myx/pkg/template/output"
)

//go:embed api.template
var APICode string

type APIValues struct {
	PipelineCode
	InputCode  input.InputCode
	OutputCode output.OutputCode
	ModelCode  model.ModelCode
}

type InputCode struct {
	Type string
	Body string
}

type PipelineCode struct {
	Pipelines   []string
	Aggregation string
}

func GenerateAPICode(values *APIValues) (string, error) {
	t, err := template.New("input").Parse(APICode)
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
