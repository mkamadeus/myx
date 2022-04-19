package model

import (
	"github.com/mkamadeus/myx/pkg/template/model/onnx"
)

type ONNXModule struct {
	Path string
}

func (module *ONNXModule) GetSessionCode() ([]string, error) {
	sessionValues := &onnx.ModelONNXBaseValues{
		Path: module.Path,
	}
	return onnx.GenerateONNXBase(sessionValues)
}

func (module *ONNXModule) GetPredictionCode() ([]string, error) {
	predictionValues := &onnx.ModelONNXPredictionValues{}
	return onnx.GenerateONNXPrediction(predictionValues)
}
