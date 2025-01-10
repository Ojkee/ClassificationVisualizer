package dataset

import "fmt"

type Point struct {
	x     float64
	y     float64
	label int
}

func NewPoint(x, y float64, label int) Point {
	return Point{
		x:     x,
		y:     y,
		label: label,
	}
}

func (p *Point) ToString() string {
	return fmt.Sprintf("(%f, %f)  :  %d", p.x, p.y, p.label)
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) GetLabel() int {
	return p.label
}
