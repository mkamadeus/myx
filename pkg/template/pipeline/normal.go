package pipeline

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed tabular_normal.template
var TabularNormalTemplate string

type TabularNormalValues struct {
	Index     int
	Name      string
	NumpyType string
}

func GenerateTabularNormalCode(values *TabularNormalValues) (string, error) {
	t, err := template.New("tabular_normal").Parse(TabularNormalTemplate)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		panic(err)
	}

	return buf.String(), nil
}
