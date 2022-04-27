package pipeline

import "github.com/mkamadeus/myx/pkg/template/pipeline/image"

type ResizeModule struct {
	Width  int
	Height int
}

func (module *ResizeModule) Run() ([]string, error) {
	values := &image.ImageResizeValues{
		Width:  module.Width,
		Height: module.Height,
	}
	return image.GenerateImageResizeCode(values)
}
