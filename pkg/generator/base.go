package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/mkamadeus/myx/pkg/spec"
	textTemplate "github.com/mkamadeus/myx/pkg/template"
)

type Generator interface {
	Name() string
	Execute(*spec.MyxSpec) error
}

func RenderSpec(s *spec.MyxSpec) error {
	// generators := make([]Generator, 0)
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
		// t, err := template.New("input").Parse(textTemplate.InputTemplate)
		// fmt.Println()
	} else {
		return fmt.Errorf("undefined input type")
	}

	// output
	t, err = template.New("output").Parse(textTemplate.OutputTemplate)
	if err != nil {
		panic(err)
	}
	buf = new(bytes.Buffer)
	err = t.Execute(buf, s.Output)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())

	// model
	t, err = template.New("model").Parse(textTemplate.ModelTemplate)
	if err != nil {
		panic(err)
	}
	buf = new(bytes.Buffer)
	err = t.Execute(buf, s.Model)
	if err != nil {
		panic(err)
	}

	// pipeline
	// for _, pipeline := range s.Pipeline {
	// 	if pipeline.Module == "preprocessing/scale" {
	// 		scalerPath := fmt.Sprintf("%v", pipeline.Metadata["path"])

	// 	}
	// }

	return nil
}
