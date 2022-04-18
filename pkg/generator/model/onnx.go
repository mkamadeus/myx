package model

import (
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/model"
	"github.com/mkamadeus/myx/pkg/template/model/onnx"
)

func RenderONNXModelCode(s *spec.MyxSpec) (*model.ModelCode, error) {
	sessionValues := &onnx.ModelONNXBaseValues{
		Path: s.Model.Path,
	}
	predictionValues := &onnx.ModelONNXPredictionValues{}
	sessionCode, err := onnx.GenerateONNXBase(sessionValues)
	if err != nil {
		return nil, err
	}
	predictionCode, err := onnx.GenerateONNXPrediction(predictionValues)
	if err != nil {
		return nil, err
	}
	return &model.ModelCode{
		Session:    sessionCode,
		Prediction: predictionCode,
	}, nil
}
