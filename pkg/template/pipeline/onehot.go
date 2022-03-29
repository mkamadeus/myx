package pipeline

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed tabular_onehot.template
var TabularOneHotTemplate string

type TabularOneHotValues struct {
	Index     int
	Name      string
	Value     string
	NumpyType string
}

func GenerateTabularOneHotCode(values *TabularOneHotValues) (string, error) {
	t, err := template.New("tabular_onehot").Parse(TabularOneHotTemplate)
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
