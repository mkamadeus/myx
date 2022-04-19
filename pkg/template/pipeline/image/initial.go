package image

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed image_initial.template
var ImageInitialTemplate string

type ImageInitialValues struct {}

func GenerateImageInitialCode(values *ImageInitialValues) ([]string, error) {
	t, err := template.New("image_resize").Parse(ImageResizeTemplate)
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