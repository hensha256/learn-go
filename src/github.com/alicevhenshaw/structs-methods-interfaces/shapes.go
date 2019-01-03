package shapes

import "math"

// Anything that implements function Area() returning float64 is classed as a 'shape'
// You don't have to say Rectangle implements interface Shape it is implicit
type Shape interface {
    Area() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base float64
	height float64
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

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}