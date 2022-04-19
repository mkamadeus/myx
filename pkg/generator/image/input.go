package input

import (
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models/spec"
	"github.com/mkamadeus/myx/pkg/template/input"
)

// TODO: implement image input rendering
func RenderImageInputSpec(s *spec.MyxSpec) (*input.InputCode, error) {
	logger.Logger.Instance.Debug("running in image input mode")
	return nil, nil
}
