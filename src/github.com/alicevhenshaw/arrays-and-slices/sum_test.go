package arrays

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	
	t.Run("collection of 5 numbers - array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		
        got := Sum(numbers)
		want := 15

		if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
	})

	t.Run("collection of 5 numbers - slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		
        got := SumSlice(numbers)
		want := 15

		if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
	})
	
	t.Run("collection of 3 numbers - slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		
		got := SumSlice(numbers)
		want := 6
		
		if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })
}

func TestSumAll(t *testing.T) {
	got := SumAll([][]int{[]int{1,2,3,4}, []int{6,7,8}, []int{9,0}})
	want := []int{10,21,9}

	if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TextSumAllTails(t *testing.T) {
	got := SumAllTails([][]int{[]int{1,2,3,4}, []int{6,7,8}, []int{9,0}, []int{4}, []int{}})
	want := []int{9,15,9,0,0}

	if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}