package model

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed model_keras_base.template
var ModelKerasBaseTemplate string

type ModelKerasBaseValues struct {
	Path string
}

//go:embed model_keras_prediction.template
var ModelKerasPredictionTemplate string

type ModelKerasPredictionValues struct{}

func GenerateKerasBase(values *ModelKerasBaseValues) (string, error) {
	t, err := template.New("onnx_model_base").Parse(ModelKerasBaseTemplate)
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

func GenerateKerasPrediction(values *ModelKerasPredictionValues) (string, error) {
	t, err := template.New("onnx_model_prediction").Parse(ModelKerasPredictionTemplate)
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
