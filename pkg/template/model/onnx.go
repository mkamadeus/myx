package model

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed model_onnx_base.template
var ModelONNXBaseTemplate string

type ModelONNXBaseValues struct {
	Path string
}

//go:embed model_onnx_prediction.template
var ModelONNXPredictionTemplate string

type ModelONNXPredictionValues struct{}

func GenerateONNXBase(values *ModelONNXBaseValues) (string, error) {
	t, err := template.New("onnx_model_base").Parse(ModelONNXBaseTemplate)
	if err != nil {
		return "", nil
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		return "", nil
	}

	return buf.String(), nil
}

func GenerateONNXPrediction(values *ModelONNXPredictionValues) (string, error) {
	t, err := template.New("onnx_model_prediction").Parse(ModelONNXPredictionTemplate)
	if err != nil {
		return "", nil
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		return "", nil
	}

	return buf.String(), nil
}
