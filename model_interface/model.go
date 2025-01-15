package model_interface

type ModelInterface interface {
	Forward(float64, float64) []float64
	Train(float64, float64, int)
	Info() string
}
