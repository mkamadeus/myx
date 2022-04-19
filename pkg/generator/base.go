package generator

import (
	"github.com/mkamadeus/myx/pkg/template/input"
	"github.com/mkamadeus/myx/pkg/template/model"
	"github.com/mkamadeus/myx/pkg/template/output"
	"github.com/mkamadeus/myx/pkg/template/pipeline"
)

type Generator interface {
	RenderSpec() (string, error)
	RenderInputSpec() (*input.InputCode, error)
	RenderOutputSpec() (*output.OutputCode, error)
	RenderModelSpec() (*model.ModelCode, error)
	RenderPipelineSpec() (*pipeline.PipelineCode, error)
}