package api

import (
	"bytes"
	_ "embed"
	"html/template"
)

//go:embed api.template
var APICode string

type APIValues struct {
	PipelineCode []string
	InputCode    string
	OutputCode   string
	ModelCode    ModelCode
}

type ModelCode struct {
	Session    string
	Prediction string
}

func GenerateAPICode() (string, error) {
	t, err := template.New("input").Parse(Model)
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
