package image

import "github.com/mkamadeus/myx/pkg/template/pipeline/image"

type InitialReadModule struct {}

func (module *InitialReadModule) Run() ([]string, error) {
	return image.GenerateImageInitialCode(&image.ImageInitialValues{})
}