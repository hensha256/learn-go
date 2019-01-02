package shapes

import "math"

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() float64 {
	x := r.width
	y := r.height
	return x + x + y + y
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}