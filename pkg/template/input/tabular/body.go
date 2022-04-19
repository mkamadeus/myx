package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed input_tabular_body.template
var TabularInputBodyTemplate string

type TabularInputBodyValues struct{}

func GenerateTabularInputBodyCode(values *TabularInputBodyValues) ([]string, error) {
	t, err := template.New("input").Parse(TabularInputBodyTemplate)
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
