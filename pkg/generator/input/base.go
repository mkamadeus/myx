package input

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/input"
)

func RenderInputSpec(s *spec.MyxSpec) (*input.InputCode, error) {
	var code *input.InputCode
	var err error

	// input
	if s.Input.Format == "tabular" {
		code, err = RenderTabularInputSpec(s)
	} else if s.Input.Format == "image" {
		// TODO: implement input spec image
		code, err = RenderImageInputSpec(s)
	} else {
		return nil, fmt.Errorf("undefined input type %s", s.Input.Format)
	}

	if err != nil {
		return nil, err
	}

	return code, nil

}
