package api

//go:embed api.template
var APICode string

type APIValues interface {
	Render() (string, error)
}
