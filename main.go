package main

import (
	"ClassificationVisualizer/dataset"
	"ClassificationVisualizer/settings"
	"ClassificationVisualizer/window"
)

func main() {
	points := dataset.GeneratePoints(50, f)
	model := Model{}
	colors := createColorsAsHex()

	main_window := window.NewWindow(
		settings.WINDOW_HEIGHT,
		settings.WINDOW_WIDHT,
		points,
		&model,
		colors,
	)
	main_window.MainLoop()
}

// TEMPLATE

// Function that model will try to fit
// (x, y) is coord of the point
// returns index of the label
func f(x, y float64) int {
	if x > y {
		return 1
	}
	return 0
}

// Custom colors, they need to match number of classes
// format RGBA
func createColorsAsHex() []uint64 {
	return []uint64{
		0xFF0000FF,
		0x0000FFFF,
	}
}

// Model has to implement this
type Model struct{}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Forward(x, y float64) []float64 {
	return []float64{1}
}

func (m *Model) Train(x, y float64, label int) {
}
