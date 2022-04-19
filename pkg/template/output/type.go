package output

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed output_type.template
var OutputTypeTemplate string

type OutputTypeValues struct {
	Name string
	Type string
}

func GenerateOutputType(values []*OutputTypeValues) ([]string, error) {
	t, err := template.New("output_type").Parse(OutputTypeTemplate)
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
