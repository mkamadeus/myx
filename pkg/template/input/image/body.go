package image

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed input_image_body.template
var ImageInputBodyTemplate string

type ImageInputBodyValues struct{}

func GenerateImageInputBodyCode(values *ImageInputBodyValues) ([]string, error) {
	t, err := template.New("input").Parse(ImageInputBodyTemplate)
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
