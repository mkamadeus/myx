package template

import _ "embed"

//go:embed input.template
var InputTemplate string

//go:embed output.template
var OutputTemplate string

//go:embed api.template
var ApiTemplate string

//go:embed model.template
var ModelTemplate string
