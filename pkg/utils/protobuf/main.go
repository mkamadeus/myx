package protobuf

import (
	"bytes"
	"text/template"
)

func RenderMainProto(contents BaseProto) (string, error) {
	t, err := template.New("baseProto").Parse(baseProtoTemplate)
	if err != nil {
		return "", err
	}

	b := bytes.Buffer{}
	if err = t.Execute(&b, contents); err != nil {
		return "", err
	}

	return b.String(), nil
}
