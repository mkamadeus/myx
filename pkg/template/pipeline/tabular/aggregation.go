package tabular

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/mkamadeus/myx/pkg/utils"
)

//go:embed tabular_aggregation.template
var TabularAggregationTemplate string

type TabularAggregationValues struct {
	Variables []string
}

func GenerateTabularAggregationCode(values *TabularAggregationValues) ([]string, error) {
	t, err := template.New("tabularaggregation").Parse(TabularAggregationTemplate)
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
