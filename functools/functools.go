package functools

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UintToRlColor(color uint64) rl.Color {
	r := uint8((color & 0xFF000000) >> (3 * 8))
	g := uint8((color & 0x00FF0000) >> (2 * 8))
	b := uint8((color & 0x0000FF00) >> (1 * 8))
	a := uint8((color & 0x000000FF) >> (0 * 8))
	return rl.NewColor(r, g, b, a)
}
