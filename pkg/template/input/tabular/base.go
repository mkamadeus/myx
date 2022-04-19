package tabular

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed input_tabular_body.template
var TabularInputBodyTemplate string

type TabularInputBodyValues struct{}

func GenerateTabularInputBodyCode(values *TabularInputBodyValues) (string, error) {
	t, err := template.New("input").Parse(TabularInputBodyTemplate)
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
