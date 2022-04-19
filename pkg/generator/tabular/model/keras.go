package model

import (
	"github.com/mkamadeus/myx/pkg/template/model/keras"
)

type KerasModule struct {
	Path string
}

func (module *KerasModule) GetSessionCode() ([]string, error) {
	sessionValues := &keras.ModelKerasBaseValues{
		Path: module.Path,
	}
	return keras.GenerateKerasBase(sessionValues)
}

func (module *KerasModule) GetPredictionCode() ([]string, error) {
	predictionValues := &keras.ModelKerasPredictionValues{}
	return keras.GenerateKerasPrediction(predictionValues)
}
