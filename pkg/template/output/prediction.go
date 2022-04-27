package output

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed output_prediction.template
var OutputPredictionTemplate string

type OutputPredictionValues struct{}

func GenerateOutputPrediction(values *OutputPredictionValues) ([]string, error) {
	t, err := template.New("output_prediction").Parse(OutputPredictionTemplate)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		return nil, err
	}

	return utils.ClearEmptyString(strings.Split(buf.String(), "\n")), nil
}
