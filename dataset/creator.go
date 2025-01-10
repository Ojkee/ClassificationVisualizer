package dataset

import "math/rand"

func GeneratePoints(n int, f func(float64, float64) int) []Point {
	retVal := make([]Point, n)
	for i := range n {
		x := rand.Float64()
		y := rand.Float64()
		label := f(x, y)
		retVal[i] = NewPoint(x, y, label)
	}
	return retVal
}
