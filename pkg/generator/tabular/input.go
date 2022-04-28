package tabular

import (
	"fmt"
	"strings"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/template/input/tabular"
)

func (g *TabularGenerator) RenderInputSpec() (*generator.InputCode, error) {

	logger.Logger.Instance.Debug("running in tabular input mode")

	values := make([]*tabular.TabularInputTypeValues, 0)
	imports := make([]string, 0)

	columns := g.Spec.Input.Metadata["columns"].([]interface{})
	for _, v := range columns {
		casted := v.(map[interface{}]interface{})
		value := &tabular.TabularInputTypeValues{
			Name: casted["name"].(string),
			Type: casted["type"].(string),
		}

		// check array type, duplicates will be thrown away later
		if strings.HasSuffix(value.Type, "[]") {
			imports = append(imports, "from typing import List")
			value.Type = fmt.Sprintf("List[%s]", models.BodyTypeMapper[value.Type[0:len(value.Type)-2]])
		} else {
			value.Type = models.BodyTypeMapper[value.Type]
		}

		values = append(values, value)
	}

	typeCode, err := tabular.GenerateTabularInputTypeCode(values)
	if err != nil {
		return nil, err
	}

	bodyCode, err := tabular.GenerateTabularInputBodyCode(&tabular.TabularInputBodyValues{})
	if err != nil {
		return nil, err
	}

	return &generator.InputCode{
		Imports: imports,
		Type:    typeCode,
		Body:    bodyCode,
	}, nil

}
