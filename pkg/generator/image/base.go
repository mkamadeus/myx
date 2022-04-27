package image

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/mkamadeus/myx/pkg/models"
)

//go:embed api.template
var APICode string

type ImageGenerator struct {
	Spec *models.MyxSpec
}

func (g *ImageGenerator) RenderCode() (string, error) {
	values, err := g.RenderAPISpec()
	if err != nil {
		return "", err
	}

	t, err := template.New("image_api_code").Parse(APICode)
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
