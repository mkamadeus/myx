package protobuf

type BaseProto struct {
	PackageName string
	Messages    []string
}

var baseProtoTemplate string = `
syntax = "protoTemplate3";

package {{.PackageName}}

{{.Messages}}
`

type MessageProto struct {
	Name   string
	Fields []string
}

var messageProtoTemplate string = `
message {{.Name}} {
	{{.Fields}}
}
`

type MessageFieldProto struct {
	VarType  string
	VarName  string
	VarOrder uint
}

var messageFieldProtoTemplate string = "{{.VarType}} {{.VarName}} = {{.VarOrder}};"
