package output

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed output_prediction.template
var OutputPredictionTemplate string

type OutputPredictionValues struct{}

func GenerateOutputPrediction(values *OutputPredictionValues) (string, error) {
	t, err := template.New("output_prediction").Parse(OutputPredictionTemplate)
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
