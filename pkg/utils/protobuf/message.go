package protobuf

import (
	"bytes"
	"text/template"
)

func RenderMessageProto(message MessageProto) (string, error) {
	t, err := template.New("message").Parse(messageProtoTemplate)
	if err != nil {
		return "", err
	}

	fieldString := ""
	for _, f := range message.Fields {
		fieldString += f
	}

	b := bytes.Buffer{}
	if err = t.Execute(&b, fieldString); err != nil {
		return "", err
	}

	return b.String(), nil
}

func RenderMessageFieldProto(field MessageFieldProto) (string, error) {
	t, err := template.New("messageField").Parse(messageFieldProtoTemplate)
	if err != nil {
		return "", err
	}

	b := bytes.Buffer{}
	if err = t.Execute(&b, field); err != nil {
		return "", err
	}

	return b.String(), nil
}
