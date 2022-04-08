package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed tabular_scaled.template
var TabularScaledTemplate string

type TabularScaledValues struct {
	Index     int
	Name      string
	Path      string
	NumpyType string
}

func GenerateTabularScaledCode(values *TabularScaledValues) ([]string, error) {
	t, err := template.New("tabular_scaled").Parse(TabularScaledTemplate)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, values)
	if err != nil {
		panic(err)
	}

	return utils.ClearEmptyString(strings.Split(buf.String(), "\n")), nil
}
