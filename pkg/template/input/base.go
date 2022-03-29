package input

import (
	_ "embed"
)

type InputCode struct {
	Type string
	Body string
}

func RenderTabularInputCode(typeValues []*TabularInputTypeValues, bodyValues *TabularInputBodyValues) (*InputCode, error) {
	typeCode, err := GenerateTabularInputTypeCode(typeValues)
	if err != nil {
		return nil, err
	}
	bodyCode, err := GenerateTabularInputBodyCode(bodyValues)
	if err != nil {
		return nil, err
	}

	return &InputCode{
		Type: typeCode,
		Body: bodyCode,
	}, nil
}
