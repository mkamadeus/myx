package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed input_tabular_type.template
var TabularInputTypeTemplate string

type TabularInputTypeValues struct {
	Name string
	Type string
}

func GenerateTabularInputTypeCode(values []*TabularInputTypeValues) ([]string, error) {
	t, err := template.New("input").Parse(TabularInputTypeTemplate)
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
