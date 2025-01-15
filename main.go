package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"

	"ClassificationVisualizer/dataset"
	"ClassificationVisualizer/layers"
	"ClassificationVisualizer/losses"
	"ClassificationVisualizer/settings"
	"ClassificationVisualizer/window"
)

func main() {
	points := dataset.GeneratePoints(1000, f)
	model := NewModel()
	colors := createColorsAsHex()

	main_window := window.NewWindow(
		settings.WINDOW_HEIGHT,
		settings.WINDOW_WIDHT,
		points,
		model,
		colors,
	)
	main_window.MainLoop()
}

// TEMPLATE

// Function that model will try to fit
// (x, y) is coord of the point
// returns index of the label
func f(x, y float64) int {
	if x < y {
		return 1
	}
	return 0
}

// Custom colors, they need to match number of classes
// format RGBA
func createColorsAsHex() []uint64 {
	return []uint64{
		0xFF0000FF,
		0xFFF8E7FF,
	}
}

// Model has to implement this
type Model struct {
	numCategories int
	fc            layers.DenseLayer
	act           layers.VSigmoid
	lr            float64
	loss          losses.MeanAbsoluteError
}

func NewModel() *Model {
	fc_ := layers.NewDenseLayer(2, 1)
	fc_.InitFilterRandom(-2, 2)
	fc_.LoadBias(&[]float64{rand.Float64()})
	return &Model{
		numCategories: 1,
		fc:            fc_,
		act:           layers.NewVSigmoid(),
		lr:            0.001,
		loss:          losses.NewMeanAbsoluteError(1, 1),
	}
}

func (m *Model) Forward(x, y float64) []float64 {
	out := m.fc.Forward(mat.NewVecDense(2, []float64{x, y}))
	activated := m.act.Forward(out)
	return activated.RawVector().Data
}

func (m *Model) Train(x, y float64, label int) {
	out := m.fc.Forward(mat.NewVecDense(2, []float64{x, y}))
	activated := m.act.Forward(out)
	// loss := m.loss.CalculateTotal(
	// 	&[]mat.VecDense{*activated},
	// 	&[]mat.VecDense{
	// 		*mat.NewVecDense(m.numCategories, functools.GetTargetVecFromLabel(label, m.numCategories)),
	// 	},
	// )
	loss := float64(label) - activated.RawVector().Data[0]
	actGrad := m.act.Backward(mat.NewVecDense(m.numCategories, []float64{loss}))
	m.fc.Backward(actGrad)
	m.fc.ApplyGrads(&m.lr, m.fc.GetOutWeightsGrads(), m.fc.GetOutBiasGrads())
}

func (m *Model) Info() string {
	return fmt.Sprintf("%v", m.fc.GetWeights())
}
