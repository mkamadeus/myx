package generator

import (
	"bytes"
	"text/template"

	"github.com/mkamadeus/myx/pkg/spec"
	textTemplate "github.com/mkamadeus/myx/pkg/template"
)

func RenderModelSpec(s *spec.MyxSpec) (string, error) {
	t, err := template.New("model").Parse(textTemplate.ModelTemplate)
	if err != nil {
		return "", nil
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, s.Model)
	if err != nil {
		return "", nil
	}

	return buf.String(), nil
}
