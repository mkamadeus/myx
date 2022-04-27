package image

import (
	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/template/input/image"
)

func (g *ImageGenerator) RenderInputSpec() (*generator.InputCode, error) {

	logger.Logger.Instance.Debug("running in image input mode")

	bodyCode, err := image.GenerateImageInputBodyCode(&image.ImageInputBodyValues{})
	if err != nil {
		return nil, err
	}

	return &generator.InputCode{
		Type: []string{},
		Body: bodyCode,
	}, nil

}
