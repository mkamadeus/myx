package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed tabular_scaled.template
var TabularScaledTemplate string

type TabularScaledValues struct {
	Index     int
	Name      string
	Target    int
	NumpyType string
}

//go:embed tabular_scaler.template
var TabularScalerTemplate string

type TabularScalerValues struct {
	Names []string
	Path  string
}

func GenerateTabularScaledCode(values *TabularScaledValues) ([]string, []string, error) {
	t, err := template.New("tabular_scaled").Parse(TabularScaledTemplate)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		return nil, err
	}

	utils.ClearEmptyString(strings.Split(buf.String(), "\n")), nil

	t, err := template.New("tabular_scaler").Parse(TabularScalerTemplate)
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

func GenerateTabularScalerCode(values *TabularScalerValues) ([]string, error) {
}
