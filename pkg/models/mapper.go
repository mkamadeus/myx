package models

var BodyTypeMapper = map[string]string{
	"float":       "float",
	"int":         "int",
	"categorical": "str",
	"string":      "str",
}

var NumpyTypeMapper = map[string]string{
	"float":       "np.float32",
	"int":         "np.int_",
	"categorical": "np.int_",
}
