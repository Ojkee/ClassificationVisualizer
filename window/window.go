package window

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"

	"ClassificationVisualizer/dataset"
	"ClassificationVisualizer/functools"
	"ClassificationVisualizer/model_interface"
	"ClassificationVisualizer/settings"
)

type Window struct {
	running bool

	height          int32
	width           int32
	backgroundColor rl.Color

	points      []dataset.Point
	labelColors []rl.Color

	model model_interface.ModelInterface
}

func NewWindow(
	height, width int32,
	points []dataset.Point,
	model model_interface.ModelInterface,
	labelColors []uint64,
) *Window {
	colors := make([]rl.Color, len(labelColors))
	for i := range len(labelColors) {
		colors[i] = functools.UintToRlColor(labelColors[i])
	}
	return &Window{
		running: true,

		height:          height,
		width:           width,
		backgroundColor: settings.DARK_GREY,

		labelColors: colors,
		points:      points,

		model: model,
	}
}

func (window *Window) MainLoop() {
	rl.InitWindow(window.width, window.height, "ClassificationVisualizer")
	rl.SetExitKey(rl.KeyNull)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for window.running {
		window.checkEvent()
		window.draw()
	}

	rl.CloseWindow()
}

func (window *Window) checkEvent() {
	window.running = !rl.WindowShouldClose()
	if rl.IsKeyPressed(rl.KeyQ) {
		window.running = false
	}
}

func (window *Window) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(window.backgroundColor)

	// window.drawGuessMesh()
	window.drawTrueLabel()

	rl.EndDrawing()
}

func (window *Window) drawTrueLabel() {
	for _, point := range window.points {
		y := int32(point.GetY() * float64(window.height))
		x := int32(point.GetX() * float64(window.width))
		color := window.labelColors[point.GetLabel()]
		rl.DrawCircle(x, y, 2, color)
	}
}

func (window *Window) drawGuessMesh() {
	for i := range window.height + 1 {
		for j := range window.width + 1 {
			x := float64(j) / float64(window.width)
			y := float64(i) / float64(window.height)
			guess := window.model.Forward(x, y)
			labelIdx, certainty := window.getLabelWithCertainty(guess)
			color := window.labelColors[labelIdx]
			color.A = uint8(certainty * 255)
			rl.DrawCircle(j, i, 8, color)
		}
	}
}

func (window *Window) getLabelWithCertainty(guess []float64) (int, float64) {
	if len(guess) == 1 {
		binaryClass := math.Round(guess[0])
		if binaryClass == 0 {
			return 0, 1 - guess[0]
		}
		return 1, guess[0]
	}
	currMaxIdx := 0
	for i := range guess {
		if guess[i] > guess[currMaxIdx] {
			currMaxIdx = i
		}
	}
	return currMaxIdx, guess[currMaxIdx]
}
