package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/mkamadeus/myx/pkg/spec"
	textTemplate "github.com/mkamadeus/myx/pkg/template"
)

func RenderSpec(s *spec.MyxSpec) error {
	// generators := make([]Generator, 0)
	var t *template.Template
	var err error
	var buf *bytes.Buffer

	// input
	inputSpec, err := RenderInputSpec(s)
	if err != nil {
		return err
	}
	fmt.Println(inputSpec)

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

	pipelineSpec, err := RenderPipelineSpec(s)
	if err != nil {
		return err
	}
	fmt.Println(pipelineSpec)

	return nil
}
