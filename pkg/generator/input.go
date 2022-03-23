package generator

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/mkamadeus/myx/pkg/spec"
	textTemplate "github.com/mkamadeus/myx/pkg/template"
)

func RenderInputSpec(s *spec.MyxSpec) (string, error) {
	var t *template.Template
	var err error
	var buf *bytes.Buffer

	// input
	if s.Input.Format == "tabular" {
		t, err = template.New("input").Parse(textTemplate.InputTemplate)
		if err != nil {
			panic(err)
		}
		buf := new(bytes.Buffer)
		err = t.Execute(buf, s.Input)
		if err != nil {
			panic(err)
		}
		fmt.Println(buf.String())
	} else if s.Input.Format == "image" {
		// TODO: implement input spec image
	} else {
		return "", fmt.Errorf("undefined input type")
	}

	return buf.String(), nil
}
