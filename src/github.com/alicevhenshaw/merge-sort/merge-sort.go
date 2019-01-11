package main

import "fmt"

func MergeSort(sliceOfNumbers []int) []int {
	length := len(sliceOfNumbers)
	sorted := make([]int, length)
	if length > 1 {
		mid := (length-1)/2 +1
		first := MergeSort(sliceOfNumbers[0:mid])
		second := MergeSort(sliceOfNumbers[mid:length])
		firstIndex := 0
		secondIndex := 0
		for i:=0; i<length; i++ {
			if secondIndex==len(second) {
				sorted[i] = first[firstIndex]
				firstIndex++
			} else if firstIndex==len(first) {
				sorted[i] = second[secondIndex]
				secondIndex++
			} else if first[firstIndex] < second[secondIndex] {
				sorted[i] = first[firstIndex]
				firstIndex++
			} else {
				sorted[i] = second[secondIndex]
				secondIndex++
			}
		}
		fmt.Println(sorted)
		return sorted
	} else {
		return sliceOfNumbers
	}
}

func main() {
	fmt.Println(MergeSort([]int{4}))
}