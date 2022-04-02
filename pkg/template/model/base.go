package model

type ModelCode struct {
	Session    string
	Prediction string
}

func RenderKerasModelCode(sessionValues *ModelKerasBaseValues, predictionValues *ModelKerasPredictionValues) (*ModelCode, error) {
	sessionCode, err := GenerateKerasBase(sessionValues)
	if err != nil {
		return nil, err
	}
	predictionCode, err := GenerateKerasPrediction(predictionValues)
	if err != nil {
		return nil, err
	}
	return &ModelCode{
		Session:    sessionCode,
		Prediction: predictionCode,
	}, nil
}

func RenderONNXModelCode(sessionValues *ModelONNXBaseValues, predictionValues *ModelONNXPredictionValues) (*ModelCode, error) {
	sessionCode, err := GenerateONNXBase(sessionValues)
	if err != nil {
		return nil, err
	}
	predictionCode, err := GenerateONNXPrediction(predictionValues)
	if err != nil {
		return nil, err
	}
	return &ModelCode{
		Session:    sessionCode,
		Prediction: predictionCode,
	}, nil
}
