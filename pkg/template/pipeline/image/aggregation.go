package image

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed image_aggregation.template
var ImageAggregationTemplate string

type ImageAggregationValues struct {
	Width   int
	Height  int
	Channel int
}

func GenerateImageAggregationCode(values *ImageAggregationValues) ([]string, error) {
	t, err := template.New("image_aggregation").Parse(ImageAggregationTemplate)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		return nil, err
	}

	return utils.ClearEmptyString(strings.Split(buf.String(), "\n")), nil
}
