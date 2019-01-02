package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
	expected := 4
	
	if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}

// Examples of how to use the code
// This runs when you add -v to your go test instruction
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}