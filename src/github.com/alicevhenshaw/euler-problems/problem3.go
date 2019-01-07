package main

import "fmt"

func ProblemThree(number int) (largestFactor int) {
	largestFactor = 1
	currentNumber := number
	currentCheck := 2
	for currentNumber > 1 {
		if currentNumber%currentCheck == 0 {
			currentNumber = currentNumber / currentCheck
			largestFactor = currentCheck
		} else {
			currentCheck++
		}
	}
	return
}