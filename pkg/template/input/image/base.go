package tabular

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed input_image_body.template
var ImageInputBodyTemplate string

type ImageInputBodyValues struct{}

func GenerateImageInputBodyCode(values *ImageInputBodyValues) (string, error) {
	t, err := template.New("input").Parse(ImageInputBodyTemplate)
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
