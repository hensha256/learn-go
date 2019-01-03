package shapes

import "testing"
import "math"

type TestCase struct {
    shape Shape
    want  float64
}

func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := rectangle.Perimeter()
	want := 40.0
	
	if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    
    areaTests := []TestCase {
        {Rectangle{12, 6}, 72},
        {Circle{10}, math.Pi * 100},
        {Triangle{4, 8}, 16},
    }
    
    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %.2f want %.2f", got, tt.want)
        }
    }

}

//Previous TestArea definition:

// func TestArea(t *testing.T) {
//     
//     checkArea := func(t *testing.T, shape Shape, want float64) {
//         t.Helper()
//         got := shape.Area()
//         if got != want {
//             t.Errorf("got %.2f want %.2f", got, want)
//         }
//     }

//     t.Run("rectangles", func(t *testing.T) {
//         rectangle := Rectangle{10.0, 50.0}
//         want := 500.0
//         checkArea(t, rectangle, want)
//     })

//     t.Run("circles", func(t *testing.T) {
//         circle := Circle{10.0}
//         want := math.Pi * 100
//         checkArea(t, circle, want)
//     })
// }