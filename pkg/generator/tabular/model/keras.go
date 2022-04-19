package model

import (
	"github.com/mkamadeus/myx/pkg/models/spec"
	"github.com/mkamadeus/myx/pkg/template/model"
	"github.com/mkamadeus/myx/pkg/template/model/keras"
)

func RenderKerasModelCode(s *spec.MyxSpec) (*model.ModelCode, error) {
	sessionValues := &keras.ModelKerasBaseValues{
		Path: s.Model.Path,
	}
	predictionValues := &keras.ModelKerasPredictionValues{}
	sessionCode, err := keras.GenerateKerasBase(sessionValues)
	if err != nil {
		return nil, err
	}
	predictionCode, err := keras.GenerateKerasPrediction(predictionValues)
	if err != nil {
		return nil, err
	}
	return &model.ModelCode{
		Session:    sessionCode,
		Prediction: predictionCode,
	}, nil
}
