package model

type ModelModule interface {
	GetSessionCode() ([]string, error)
	GetPredictionCode() ([]string, error)
}
