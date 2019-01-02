package shapes

import "testing"
import "math"

func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := rectangle.Perimeter()
	want := 40.0
	
	if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{10.0, 50.0}
        got := rectangle.Area()
        want := 500.0

        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10.0}
        got := circle.Area()
        want := math.Pi * 100

        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    })
}