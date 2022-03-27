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

//go:embed model_onnx_base.template
var ModelONNXBaseTemplate string

type ModelONNXBaseValues struct {
	Path string
}

//go:embed model_onnx_prediction.template
var ModelONNXPredictionTemplate string

type ModelKerasPredictionValues struct{}

type ModelONNXPredictionValues struct {
	Pipelines []string
}

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
