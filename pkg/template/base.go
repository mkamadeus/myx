package template

type TemplateWriter interface {
	Generate() ([]string, error)
}
