package functools

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func UintToRlColor(color uint64) rl.Color {
	r := uint8((color & 0xFF000000) >> (3 * 8))
	g := uint8((color & 0x00FF0000) >> (2 * 8))
	b := uint8((color & 0x0000FF00) >> (1 * 8))
	a := uint8((color & 0x000000FF) >> (0 * 8))
	return rl.NewColor(r, g, b, a)
}

func GetTargetVecFromLabel(label, numCategories int) []float64 {
	if numCategories == 1 {
		return []float64{float64(label)}
	}
	retVal := make([]float64, numCategories)
	retVal[label] = 1
	return retVal
}

func GetLabelWithCertainty(guess []float64) (int, float64) {
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
