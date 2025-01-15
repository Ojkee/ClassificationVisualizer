package window

import (
	"fmt"

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
	pointIdx    int
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

		points:      points,
		pointIdx:    0,
		labelColors: colors,

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
		window.trainStep()
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

	window.drawGuessMesh()
	window.drawTrueLabel()

	rl.EndDrawing()
}

func (window *Window) drawTrueLabel() {
	for _, point := range window.points {
		y := int32(point.GetY() * float64(window.height))
		x := int32(point.GetX() * float64(window.width))
		color := window.labelColors[point.GetLabel()]
		rl.DrawCircle(x, y, settings.POINTS_D, color)
	}
}

func (window *Window) drawGuessMesh() {
	for i := 0; i <= int(window.height); i += settings.MESH_GAP_Y {
		for j := 0; j <= int(window.height); j += settings.MESH_GAP_Y {
			x := float64(j) / float64(window.width)
			y := float64(i) / float64(window.height)
			guess := window.model.Forward(x, y)
			labelIdx, certainty := functools.GetLabelWithCertainty(guess)
			color := window.labelColors[labelIdx]
			color.A = uint8(certainty * 255)
			rl.DrawCircle(int32(j), int32(i), 1, color)
		}
	}
}

func (window *Window) trainStep() {
	if window.pointIdx >= len(window.points) {
		return
	}
	p := window.points[window.pointIdx]
	x := p.GetX()
	y := p.GetY()
	label := p.GetLabel()
	guess := window.model.Forward(x, y)
	labelIdx, _ := functools.GetLabelWithCertainty(guess)
	fmt.Println()
	fmt.Printf("%s\n", window.model.Info())
	fmt.Printf("(%f, %f)\n\tlabel: %d\n\tpredi: %d\n\tguess: %f\n", x, y, label, labelIdx, guess[0])
	window.model.Train(x, y, label)
	window.pointIdx += 1
}
