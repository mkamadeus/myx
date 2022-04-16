package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed tabular_labeller.template
var TabularLabellerTemplate string

type TabularLabellerValues struct {
	Names []string
	Path  string
}

//go:embed tabular_label.template
var TabularLabelTemplate string

type TabularLabelValues struct {
	Index     int
	Name      string
	Position     int
}

func GenerateTabularLabelCode(values *TabularLabelValues) ([]string, error) {
	t, err := template.New("tabular_label").Parse(TabularLabelTemplate)
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


func GenerateTabularLabellerCode(values *TabularLabellerValues) ([]string, error) {
	t, err := template.New("tabular_labeller").Parse(TabularLabellerTemplate)
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

