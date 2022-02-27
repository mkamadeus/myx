package spec

type InputSpec struct {
	Format   string                   `yaml:"format"`
	Metadata []map[string]interface{} `yaml:"meta"`
}
