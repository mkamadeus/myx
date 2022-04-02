package output

type OutputCode struct {
	Type       string
	Prediction string
}

func RenderOutputCode(typeValues []*OutputTypeValues, predictionValues *OutputPredictionValues) (*OutputCode, error) {
	typeCode, err := GenerateOutputType(typeValues)
	if err != nil {
		return nil, err
	}
	predictionCode, err := GenerateOutputPrediction(predictionValues)
	if err != nil {
		return nil, err
	}

	return &OutputCode{
		Type:       typeCode,
		Prediction: predictionCode,
	}, nil
}
