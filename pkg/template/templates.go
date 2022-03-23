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

//go:embed tabular_normal.template
var TabularNormalTemplate string

//go:embed tabular_scaled.template
var TabularScaledTemplate string

//go:embed tabular_onehot.template
var TabularOnehotTemplate string
