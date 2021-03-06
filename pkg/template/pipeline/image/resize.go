package image

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed image_resize.template
var ImageResizeTemplate string

type ImageResizeValues struct {
	Height int
	Width  int
}

func GenerateImageResizeCode(values *ImageResizeValues) ([]string, error) {
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
