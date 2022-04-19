package tabular

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed input_tabular_type.template
var TabularInputTypeTemplate string

type TabularInputTypeValues struct {
	Name string
	Type string
}

func GenerateTabularInputTypeCode(values []*TabularInputTypeValues) (string, error) {
	t, err := template.New("input").Parse(TabularInputTypeTemplate)
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
