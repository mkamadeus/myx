package generator

import (
	"bytes"
	"text/template"

	"github.com/mkamadeus/myx/pkg/spec"
	textTemplate "github.com/mkamadeus/myx/pkg/template"
)

func RenderOutputSpec(s *spec.MyxSpec) (string, error) {
	t, err := template.New("output").Parse(textTemplate.OutputTemplate)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, s.Output)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
