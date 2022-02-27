package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/mkamadeus/myx/pkg/spec"
)

type Generator interface {
	Name() string
	Execute(*spec.MyxSpec) error
}

//go:embed input.py.template
var inputTemplate string

func RenderSpec(s *spec.MyxSpec) error {
	// generators := make([]Generator, 0)

	// check input
	if s.Input.Format == "tabular" {
		fmt.Println(inputTemplate)
		t, err := template.New("input").Parse(inputTemplate)
		if err != nil {
			panic(err)
		}
		buf := new(bytes.Buffer)
		err = t.Execute(buf, s.Input)
		if err != nil {
			panic(err)
		}
		fmt.Println(buf.String())

		// generators = append(generators, input.TabularGenerator{})
	}

	// for _, g := range generators {
	// 	err := g.Execute(s)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
