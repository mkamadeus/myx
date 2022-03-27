package pipeline

import (
	"bytes"
	"text/template"
)

//go:embed tabular_normal.template
var TabularNormalTemplate string

type TabularNormalValues struct {
	Index     int
	Name      string
	NumpyType string
}

//go:embed tabular_scaled.template
var TabularScaledTemplate string

type TabularScaledValues struct {
	Index     int
	Name      string
	Path      string
	NumpyType string
}

//go:embed tabular_onehot.template
var TabularOneHotTemplate string

type TabularOneHotValues struct {
	Index     int
	Name      string
	Value     string
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

func GenerateTabularScaledCode(values *TabularScaledValues) (string, error) {
	t, err := template.New("tabular_scaled").Parse(TabularScaledTemplate)
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
