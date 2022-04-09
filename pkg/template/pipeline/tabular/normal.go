package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed tabular_normal.template
var TabularNormalTemplate string

type TabularNormalValues struct {
	Index     int
	Name      string
	NumpyType string
}

func GenerateTabularNormalCode(values *TabularNormalValues) ([]string, error) {
	t, err := template.New("tabular_normal").Parse(TabularNormalTemplate)
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
