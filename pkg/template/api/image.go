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

type ImageAPIValues struct {
	PipelineCode *pipeline.PipelineCode
	InputCode    *input.InputCode
	OutputCode   *output.OutputCode
	ModelCode    *model.ModelCode
}

func (values *ImageAPIValues) Render() (string, error) {
	t, err := template.New("image_api_code").Parse(APICode)
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
