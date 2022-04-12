package pipeline

import (
	"fmt"

	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/mkamadeus/myx/pkg/template/pipeline"
)

func RenderPipelineSpec(s *spec.MyxSpec) (*pipeline.PipelineCode, error) {

	// pipeline
	if s.Input.Format == "tabular" {

	} else if s.Input.Format == "image" {
		return nil, nil
	}

	return nil, fmt.Errorf("invalid input format")
}
