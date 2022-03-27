package output

import (
	"bytes"
	"text/template"
)

//go:embed output.template
var OutputTemplate string

type OutputValues []struct {
	Name string
	Type string
}

func GenerateOutputCode(values *OutputValues) (string, error) {
	t, err := template.New("tabular_normal").Parse(OutputTemplate)
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
