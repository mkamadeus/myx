package output

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed output_type.template
var OutputTypeTemplate string

type OutputTypeValues struct {
	Name string
	Type string
}

func GenerateOutputType(values []*OutputTypeValues) (string, error) {
	t, err := template.New("output_type").Parse(OutputTypeTemplate)
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
