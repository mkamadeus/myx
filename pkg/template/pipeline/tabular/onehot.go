package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed tabular_onehot.template
var TabularOneHotTemplate string

type TabularOneHotValues struct {
	Index     int
	Name      string
	Value     string
}

func GenerateTabularOneHotCode(values *TabularOneHotValues) ([]string, error) {
	t, err := template.New("tabular_onehot").Parse(TabularOneHotTemplate)
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
