package generator

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	textTemplate "github.com/mkamadeus/myx/pkg/template"
)

func RenderInputSpec(s *spec.MyxSpec) (string, error) {
	// input
	if s.Input.Format == "tabular" {
		logger.Logger.Instance.Debug("running in tabular input mode")
		t, err := template.New("input").Parse(textTemplate.InputTemplate)
		if err != nil {
			panic(err)
		}
		buf := new(bytes.Buffer)
		err = t.Execute(buf, s.Input)
		if err != nil {
			panic(err)
		}
		return buf.String(), nil
	} else if s.Input.Format == "image" {
		// TODO: implement input spec image
	}

	return "", fmt.Errorf("undefined input type %s", s.Input.Format)
}
