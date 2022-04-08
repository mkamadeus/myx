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
	NumpyType string
}

func GenerateTabularOneHotCode(values *TabularOneHotValues) ([]string, error) {
	t, err := template.New("tabular_onehot").Parse(TabularOneHotTemplate)
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
