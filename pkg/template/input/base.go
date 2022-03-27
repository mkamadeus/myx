package input

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed input.template
var InputTemplate string

type InputValues struct {
	Name string
	Type string
}

func GenerateInputCode(values []*InputValues) (string, error) {
	t, err := template.New("input").Parse(InputTemplate)
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
